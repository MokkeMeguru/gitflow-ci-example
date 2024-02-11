# gitflow-ci-example
example for git-flow ci


```mermaid
%%{init: { 'gitGraph': { 'mainBranchName': 'main'}} }%%
gitGraph
  commit id: "initialize repo"
  branch develop order:3
  commit id: "initialize develop branch"
  branch feature_1 order:4
  checkout feature_1
  commit id: "add something_1"
  commit id: "fix something"
  checkout develop
  branch feature_2 order:5
  checkout feature_2
  commit id: "add something_2"
  checkout main
  branch hotfix order:1
  checkout hotfix
  commit id: "hotfix"
  checkout main
  merge hotfix
  checkout develop
  merge hotfix
  checkout develop
  merge feature_1
  merge feature_2
  branch release order:2
  checkout release
  commit id: "release"
  commit id: "tiny fix on release"
  checkout main
  merge release
  checkout release
  commit id: "release to develop"
  checkout develop
  merge release
```

## status

- GitFlow
  - [x] Release
    - [x] Release to Main
    - [x] Release to Develop (after release to main PR merged into main branch)
  - [ ] Hotfix
- [ ] Workflow Run
