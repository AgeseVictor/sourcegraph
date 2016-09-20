package localstore

import (
	"database/sql"
	"time"

	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"sourcegraph.com/sourcegraph/sourcegraph/api/sourcegraph"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/dbutil"
	"sourcegraph.com/sourcegraph/sourcegraph/pkg/store"
	"sourcegraph.com/sourcegraph/sourcegraph/services/backend/accesscontrol"
)

func init() {
	AppSchema.Map.AddTableWithName(dbUser{}, "users").SetKeys(true, "UID")
	AppSchema.CreateSQL = append(AppSchema.CreateSQL,
		"ALTER TABLE users ALTER COLUMN login TYPE citext",
		"CREATE UNIQUE INDEX users_login ON users(login)",
		`ALTER TABLE users ALTER COLUMN registered_at TYPE timestamp with time zone USING registered_at::timestamp with time zone;`,
		`CREATE INDEX users_login_ci ON users((lower(login)) text_pattern_ops);`,
		`ALTER TABLE users ALTER COLUMN betas TYPE text ARRAY USING betas::text[]`,
	)
}

// dbUser DB-maps a sourcegraph.User object.
type dbUser struct {
	UID            int
	Login          string
	Name           string
	IsOrganization bool
	AvatarURL      string `db:"avatar_url"`
	Location       string
	Company        string
	HomepageURL    string `db:"homepage_url"`
	Disabled       bool   `db:"disabled"`
	Write          bool
	Admin          bool
	Betas          *dbutil.StringSlice
	BetaRegistered bool       `db:"beta_registered"`
	RegisteredAt   *time.Time `db:"registered_at"`
}

func (u *dbUser) toUser() *sourcegraph.User {
	var betas []string
	if u.Betas != nil {
		betas = u.Betas.Slice
	}
	return &sourcegraph.User{
		UID:            int32(u.UID),
		Login:          u.Login,
		Name:           u.Name,
		IsOrganization: u.IsOrganization,
		AvatarURL:      u.AvatarURL,
		Location:       u.Location,
		Company:        u.Company,
		HomepageURL:    u.HomepageURL,
		Disabled:       u.Disabled,
		Write:          u.Write,
		Admin:          u.Admin,
		Betas:          betas,
		BetaRegistered: u.BetaRegistered,
		RegisteredAt:   ts(u.RegisteredAt),
	}
}

func (u *dbUser) fromUser(u2 *sourcegraph.User) {
	u.UID = int(u2.UID)
	u.Login = u2.Login
	u.Name = u2.Name
	u.IsOrganization = u2.IsOrganization
	u.AvatarURL = u2.AvatarURL
	u.Location = u2.Location
	u.Company = u2.Company
	u.HomepageURL = u2.HomepageURL
	u.Disabled = u2.Disabled
	u.Write = u2.Write
	u.Admin = u2.Admin
	if len(u2.Betas) > 0 {
		u.Betas = &dbutil.StringSlice{Slice: u2.Betas}
	}
	u.BetaRegistered = u2.BetaRegistered
	u.RegisteredAt = tm(u2.RegisteredAt)
}

func toUsers(us []*dbUser) []*sourcegraph.User {
	u2s := make([]*sourcegraph.User, len(us))
	for i, u := range us {
		u2s[i] = u.toUser()
	}
	return u2s
}

// users is a DB-backed implementation of the Users store.
type users struct{}

var _ store.Users = (*users)(nil)

func (s *users) Get(ctx context.Context, userSpec sourcegraph.UserSpec) (*sourcegraph.User, error) {
	if err := accesscontrol.VerifyUserHasReadAccess(ctx, "Users.Get", nil); err != nil {
		return nil, err
	}
	if userSpec.UID == 0 {
		return nil, &store.UserNotFoundError{}
	}
	return s.getByUID(ctx, int(userSpec.UID))
}

func (s *users) GetWithLogin(ctx context.Context, login string) (*sourcegraph.User, error) {
	if err := accesscontrol.VerifyUserHasReadAccess(ctx, "Users.GetWithLogin", nil); err != nil {
		return nil, err
	}
	return s.getByLogin(ctx, login)
}

// getByUID returns the user with the given uid, if such a user
// exists in the database.
func (s *users) getByUID(ctx context.Context, uid int) (*sourcegraph.User, error) {
	user, err := s.getBySQL(ctx, "uid=$1", uid)
	if err == sql.ErrNoRows {
		err = &store.UserNotFoundError{UID: uid}
	}
	return user, err
}

// getByLogin returns the user with the given login, if such a user
// exists in the database.
func (s *users) getByLogin(ctx context.Context, login string) (*sourcegraph.User, error) {
	user, err := s.getBySQL(ctx, "login=$1", login)
	if err == sql.ErrNoRows {
		err = &store.UserNotFoundError{Login: login}
	}
	return user, err
}

// getBySQL returns a user matching the SQL query (if any exists). A
// "LIMIT 1" clause is appended to the query before it is executed.
func (s *users) getBySQL(ctx context.Context, query string, args ...interface{}) (*sourcegraph.User, error) {
	var user dbUser
	if err := appDBH(ctx).SelectOne(&user, "SELECT * FROM users WHERE ("+query+") LIMIT 1", args...); err != nil {
		return nil, err
	}
	return user.toUser(), nil
}

func (s *users) GetUIDByGitHubID(ctx context.Context, githubUID int) (int32, error) {
	if err := accesscontrol.VerifyUserHasReadAccess(ctx, "Users.GetUIDByGitHubID", nil); err != nil {
		return 0, err
	}
	uid, err := appDBH(ctx).SelectInt(`SELECT "user" FROM ext_auth_token WHERE host='github.com' AND (NOT disabled) AND ext_uid=$1;`, githubUID)
	if err == sql.ErrNoRows || uid == 0 {
		err = grpc.Errorf(codes.NotFound, "no external auth token for github user %d", githubUID)
	}
	return int32(uid), err
}
