#!/usr/bin/env zsh

. $(which ansi_colors)
. $(which git_utils)

function refresh() {
    REPO_NAME=${PWD##*/}
    REPO_URL=$(repo_url)
    VERSION=$(version)
    NEW_TIMESTAMP=$(gdate '+%_s+%N') # standard macOS `date` does not work with 'N'
    NEW_TAG="${VERSION}-${NEW_TIMESTAMP}"
    MAJOR=$(major)
    MINOR=$(minor)
    BRANCH=$(current_branch)

    # cd to git repo root
    cd "$(git rev-parse --show-toplevel || echo .)"


	hr
	side "REF --- REPOSITORY REFRESH"
	br

	canary "Repo: $REPO_NAME  version $VERSION"
    canary "Directory: $PWD"
	canary "URL: $REPO_URL"
    canary "New Version: $NEW_TAG"

	hr
	br

    if (( $(major) == 0 )); then
        cherry "Major version is $MAJOR ... dev autoupdates are enabled."
    fi
	side  "go build and mod tidy"
	go mod tidy >/dev/null 2>&1 && go mod verify >/dev/null 2>&1
	br

	side "go doc update"
	go doc >| go.doc
    cat go.doc | tail -n 5
	br

	side "git add all"
	git add --all
	git status | tail -n 5
	br

    printf -v commit_list '  %b\n'  "$(git status -s)"
    printf -v commit_message '%b\n' "GoBot [$NEW_TAG]: dev (pre v1.0) autosave - progress and formatting${commit_list}"

    git tag $NEW_TAG
	side "git commit -m $commit_message"
    side "$commit_list"
    git commit -m "$commit_message"
	br

	side "git push origin $BRANCH"
	git push origin $BRANCH
    git push origin --tags
	br
	hr
    cd -
}

# prefer functions repo_url and version


refresh
