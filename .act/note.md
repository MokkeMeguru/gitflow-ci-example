## Cheetsheet

```
gh act -P ubuntu-latest=nektos/act-environments-ubuntu:18.04 -W .github/workflows/create-release.yml --job create_release_pr -e .act/create-release-pr.event.json -s GITHUB_TOKEN="$(gh auth token)" --container-architecture linux/amd64
```
