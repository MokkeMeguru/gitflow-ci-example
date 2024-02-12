# Git Flow Test Senario

## Release Flow

1. Create Feature Branch to Develop Branch (feature/xxx)
2. Merge Feture Branch
3. Create Release PR using workflow dispatch
  - check the PR's title and body
4. Create to Release PR (to-release/xxx) And Merge into Release Branch
  - check Release PR's additional comments
5. Merge Release PR
  - check the opened Release to Develop PR for the Release PR
