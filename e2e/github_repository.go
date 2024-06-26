package test

var testCasesGitHubRepository = []testCase{
	{
		path:         "data.repository.code_review_not_required",
		failedEntity: "bad_branch_protection",
		passedEntity: "good_branch_protection",
	},
	{
		path:         "data.repository.code_review_by_two_members_not_required",
		failedEntity: "bad_branch_protection",
		passedEntity: "good_branch_protection",
	},
	{
		path:         "data.repository.missing_default_branch_protection",
		failedEntity: "no_branch_protection",
		passedEntity: "good_branch_protection",
	},
	{
		path:         "data.repository.missing_default_branch_protection_deletion",
		failedEntity: "bad_branch_protection",
		passedEntity: "good_branch_protection",
	},
	{
		path:         "data.repository.missing_default_branch_protection_force_push",
		failedEntity: "bad_branch_protection",
		passedEntity: "good_branch_protection",
	},
	{
		path:         "data.repository.non_linear_history",
		failedEntity: "bad_branch_protection",
		passedEntity: "good_branch_protection",
	},
	{
		path:         "data.repository.pushes_are_not_restricted",
		failedEntity: "bad_branch_protection",
		passedEntity: "good_branch_protection",
	},
	{
		path:         "data.repository.requires_branches_up_to_date_before_merge",
		failedEntity: "bad_branch_protection",
		passedEntity: "good_branch_protection",
	},
	{
		path:         "data.repository.requires_status_checks",
		failedEntity: "bad_branch_protection",
		passedEntity: "good_branch_protection",
	},
	{
		path:         "data.repository.code_review_not_limited_to_code_owners",
		failedEntity: "bad_branch_protection",
		passedEntity: "good_branch_protection",
	},
	{
		path:         "data.repository.dismisses_stale_reviews",
		failedEntity: "bad_branch_protection",
		passedEntity: "good_branch_protection",
	},
	{
		path:         "data.repository.no_signed_commits",
		failedEntity: "bad_branch_protection",
		passedEntity: "good_branch_protection",
	},
	{
		path:         "data.repository.no_conversation_resolution",
		failedEntity: "bad_branch_protection",
		passedEntity: "good_branch_protection",
	},
	{
		path:         "data.repository.review_dismissal_allowed",
		failedEntity: "bad_branch_protection",
		passedEntity: "good_branch_protection",
	},
	{
		path:         "data.repository.forking_allowed_for_repository",
		failedEntity: "bad_repo",
		passedEntity: "good_repo",
	},
	{
		path:         "data.repository.secret_scanning_not_enabled",
		failedEntity: "bad_public_repo",
		passedEntity: "good_public_repo",
	},
}
