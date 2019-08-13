export FORCE_COLOR = true

PRE_COMMIT_HOOK_PATH = .git/hooks/pre-commit
PRE_COMMIT_TEMPLATE_HOOK_PATH = hooks/pre-commit.sh

.PHONY: bootstrap test

test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

bootstrap:
	make check-pre-commit-hook
	cp $(PRE_COMMIT_TEMPLATE_HOOK_PATH) $(PRE_COMMIT_HOOK_PATH)
	chmod +x $(PRE_COMMIT_HOOK_PATH)

check-pre-commit-hook:
ifneq ("$(wildcard $(PRE_COMMIT_HOOK_PATH))", "")
	echo "Pre-commit hook already exists, reseting"
	rm $(PRE_COMMIT_HOOK_PATH)
else
	echo "Pre-commit hook is missing"
endif
