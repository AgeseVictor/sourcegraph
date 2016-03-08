package app

import (
	"encoding/json"
	"fmt"
	htmpl "html/template"
	"net/url"
	"path"
	"strings"
	"time"

	"golang.org/x/net/context"

	"sourcegraph.com/sourcegraph/appdash"
	"src.sourcegraph.com/sourcegraph/go-sourcegraph/sourcegraph"
	"src.sourcegraph.com/sourcegraph/server/accesscontrol"

	"src.sourcegraph.com/sourcegraph/app/appconf"
	"src.sourcegraph.com/sourcegraph/app/assets"
	"src.sourcegraph.com/sourcegraph/app/internal/schemautil"
	"src.sourcegraph.com/sourcegraph/app/internal/tmpl"
	"src.sourcegraph.com/sourcegraph/app/router"
	"src.sourcegraph.com/sourcegraph/auth/authutil"
	"src.sourcegraph.com/sourcegraph/auth/idkey"
	"src.sourcegraph.com/sourcegraph/conf"
	"src.sourcegraph.com/sourcegraph/platform"
	"src.sourcegraph.com/sourcegraph/sgx/buildvar"
	"src.sourcegraph.com/sourcegraph/sourcecode"
	"src.sourcegraph.com/sourcegraph/util/envutil"
	"src.sourcegraph.com/sourcegraph/util/handlerutil"
	"src.sourcegraph.com/sourcegraph/util/metricutil"
	"src.sourcegraph.com/sourcegraph/util/textutil"
	"src.sourcegraph.com/sourcegraph/util/timeutil"
	"src.sourcegraph.com/sourcegraph/util/traceutil/appdashctx"
)

func init() {
	tmpl.FuncMap = tmplFuncs
}

var tmplFuncs = htmpl.FuncMap{
	"personLabel": personLabel,

	"repoBasename": repoBasename,
	"repoLink":     repoLink,

	"repoMetaDescription": repoMetaDescription,

	"defQualifiedName":            sourcecode.DefQualifiedName,
	"defQualifiedNameAndType":     sourcecode.DefQualifiedNameAndType,
	"overrideStyleViaRegexpFlags": sourcecode.OverrideStyleViaRegexpFlags,

	"appconf":   func() interface{} { return &appconf.Flags },
	"authFlags": func() *authutil.Flags { return &authutil.ActiveFlags },

	"buildClass":  buildClass,
	"buildStatus": buildStatus,

	"add": func(a, b int) int { return a + b },
	"json": func(v interface{}) (string, error) {
		b, err := json.Marshal(v)
		if err != nil {
			return "", err
		}
		return string(b), nil
	},

	"customLogo":         func() htmpl.HTML { return appconf.Flags.CustomLogo },
	"customFeedbackForm": func() htmpl.HTML { return appconf.Flags.CustomFeedbackForm },
	"uiBuild":            func() bool { return !appconf.Flags.NoUIBuild },

	"urlTo":                router.Rel.URLTo,
	"urlToUserSubroute":    router.Rel.URLToUserSubroute,
	"urlToRepo":            router.Rel.URLToRepo,
	"urlToRepoRev":         router.Rel.URLToRepoRev,
	"urlToRepoBuild":       router.Rel.URLToRepoBuild,
	"urlToRepoSubroute":    router.Rel.URLToRepoSubroute,
	"urlToRepoSubrouteRev": router.Rel.URLToRepoSubrouteRev,
	"urlToRepoTreeEntry":   router.Rel.URLToRepoTreeEntry,
	"urlToRepoCommit":      router.Rel.URLToRepoCommit,
	"urlToRepoApp":         router.Rel.URLToRepoApp,
	"urlWithSchema":        schemautil.URLWithSchema,
	"urlToDef":             router.Rel.URLToDef,
	"urlToDefAtRev":        router.Rel.URLToDefAtRev,
	"urlToDefSubroute":     router.Rel.URLToDefSubroute,
	"urlToWithReturnTo":    urlToWithReturnTo,

	"fileToBreadcrumb":      FileToBreadcrumb,
	"fileLinesToBreadcrumb": FileLinesToBreadcrumb,
	"snippetToBreadcrumb":   SnippetToBreadcrumb,
	"router":                func() *router.Router { return router.Rel },

	"flattenName":     handlerutil.FlattenName,
	"flattenNameHTML": handlerutil.FlattenNameHTML,

	"schemaMatchesExceptListAndSortOptions": schemautil.SchemaMatchesExceptListAndSortOptions,

	"classForRoute": func(route string) string {
		parts := strings.Split(route, ".")
		classes := make([]string, len(parts))
		for i := range parts {
			classes[i] = "route-" + strings.Join(parts[:i+1], "-")
		}
		return strings.Join(classes, " ")
	},

	"ifTrue": func(cond bool, v interface{}) interface{} {
		if cond {
			return v
		}
		return nil
	},

	"commitSummary":       commitSummary,
	"commitRestOfMessage": commitRestOfMessage,

	"toString2":    func(v interface{}) string { return fmt.Sprintf("%s", v) },
	"sanitizeHTML": sanitizeHTML,
	"textFromHTML": textutil.TextFromHTML,
	"timeOrNil":    timeutil.TimeOrNil,
	"timeAgo":      timeutil.TimeAgo,
	"now":          time.Now,
	"duration":     duration,
	"isNil":        isNil,
	"minTime":      minTime,
	"pathJoin":     path.Join,
	"toInt": func(v interface{}) (int, error) {
		switch v := v.(type) {
		case int:
			return v, nil
		case uint32:
			return int(v), nil
		case int32:
			return int(v), nil
		case uint:
			return int(v), nil
		case uint64:
			return int(v), nil
		case int64:
			return int(v), nil
		}
		return 0, fmt.Errorf("toInt: unexpected type %T", v)
	},

	"truncate":         textutil.Truncate,
	"truncateCommitID": truncateCommitID,
	"maxLen":           maxLen,
	"displayURL": func(urlStr string) string {
		return strings.TrimSuffix(strings.TrimPrefix(strings.TrimPrefix(urlStr, "https://"), "http://"), "/")
	},

	"assetURL": assets.URL,

	"getClientIDOrHostName": func(ctx context.Context, appURL *url.URL) string {
		clientID := idkey.FromContext(ctx).ID
		if clientID != "" {
			// return a shortened clientID, to match the clientID logged
			// in eventsutil/events.go:getShortClientID.
			if len(clientID) > 6 {
				return clientID[:6]
			}
			return clientID
		}
		if appURL == nil {
			return "unknown-host"
		}
		return appURL.Host
	},

	"hasField": hasStructField,

	"ifTemplate":                ifTemplate,
	"googleAnalyticsTrackingID": func() string { return appconf.Flags.GoogleAnalyticsTrackingID },
	"heapAnalyticsID":           func() string { return appconf.Flags.HeapAnalyticsID },

	"deployedGitCommitID": func() string { return envutil.GitCommitID },
	"hostname":            func() string { return hostname },

	"showRepoRevSwitcher": showRepoRevSwitcher,

	"orderedRepoEnabledFrames": func(repo *sourcegraph.Repo, repoConf *sourcegraph.RepoConfig) []platform.RepoFrame {
		frames, orderedIDs := orderedRepoEnabledFrames(repo, repoConf)
		orderedFrames := make([]platform.RepoFrame, len(orderedIDs))
		for i, id := range orderedIDs {
			orderedFrames[i] = frames[id]
		}
		return orderedFrames
	},
	"platformSearchFrames": func() map[string]platform.SearchFrame {
		return platform.SearchFrames()
	},
	"showSearchForm":     showSearchForm,
	"fileSearchDisabled": func() bool { return appconf.Flags.DisableSearch },
	"disableCloneURL":    func() bool { return appconf.Flags.DisableCloneURL },

	"isAdmin": func(ctx context.Context, method string) bool {
		return accesscontrol.VerifyUserHasAdminAccess(ctx, method) == nil
	},

	"activeRepoApp": func(currentURL *url.URL, repoURI, appID string) (bool, error) {
		u, err := router.Rel.URLToRepoApp(repoURI, appID)
		if err != nil {
			return false, err
		}
		return strings.HasPrefix(currentURL.Path, u.Path), nil
	},

	"publicRavenDSN": func() string { return conf.PublicRavenDSN },

	"urlToAppdashTrace": func(ctx context.Context, trace appdash.ID) *url.URL {
		return appdashctx.AppdashURL(ctx).ResolveReference(&url.URL{
			Path: fmt.Sprintf("/traces/%v", trace),
		})
	},

	"buildvar":        func() buildvar.Vars { return buildvar.All },
	"updateAvailable": updateAvailable,

	"showDataCollectionMessage": metricutil.EnableForwarding,
}
