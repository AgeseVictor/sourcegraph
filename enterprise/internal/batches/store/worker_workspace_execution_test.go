	"github.com/cockroachdb/errors"

	workspace := &btypes.BatchSpecWorkspace{BatchSpecID: batchSpec.ID, RepoID: repo.ID, Steps: []batcheslib.Step{}}
	executionStore := &batchSpecWorkspaceExecutionWorkerStore{Store: workStore, observationContext: &observation.TestContext}
		tokenID, _, err := database.AccessTokens(db).CreateInternal(ctx, user.ID, []string{"user:all"}, "testing", user.ID)
		_, err = database.AccessTokens(db).GetByID(ctx, tokenID)
		database.Mocks.AccessTokens.HardDeleteByID = func(id int64) error {
		defer func() { database.Mocks.AccessTokens.HardDeleteByID = nil }()
	workspace := &btypes.BatchSpecWorkspace{BatchSpecID: batchSpec.ID, RepoID: repo.ID, Steps: []batcheslib.Step{}}
	executionStore := &batchSpecWorkspaceExecutionWorkerStore{Store: workStore, observationContext: &observation.TestContext}
		tokenID, _, err := database.AccessTokens(db).CreateInternal(ctx, user.ID, []string{"user:all"}, "testing", user.ID)
	_, err = database.AccessTokens(db).GetByID(ctx, tokenID)