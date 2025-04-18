# Maple Pie
This project is intended to provide visualizations of open election data from Canadian governments.

The name Maple Pie, a Canadian dessert, comes from the association of maple with Canada and from pie charts. A website providing graphs of Canadian data is the vision for this project.

The project is in its infancy. The intent at this time is to construct commands to handle API calls to various endpoints. The code is written in Golang.

# For developers

[![semantic-release: angular](https://img.shields.io/badge/semantic--release-angular-e10079?logo=semantic-release)](https://github.com/semantic-release/semantic-release)

## Repository structure

This repository is structured following [the "Organizing a Go module" guide for a Server Project][go-org-guide-server].

[go-org-guide-server]: https://go.dev/doc/modules/layout#server-project

## Versioning & commit message conventions

Automated Versioning and Release occurs exclusively on the `main` branch and is handled by the GitHub Actions CI pipeline.

To increase the version, commit message headers must follow [`semantic-release`][semantic-release] conventions, which in turn follow [Angular Commit Message Conventions][angular-conv]. This convention is partially enforced by the GitHub Actions CI pipeline.

What is not enforced is the free-form change descriptions. Please self-enforce these rules:
- use the imperative, present tense: ex. "change" not "changed" nor "changes" nor "changing"
- don't capitalize first letter
- no dot (.) at the end

In the near future, an issue tracker may be introduced to be referenced in the footer of the commit messages.

[semantic-release]: https://github.com/semantic-release/semantic-release
[angular-conv]: https://github.com/angular/angular.js/blob/master/DEVELOPERS.md#-git-commit-guidelines
