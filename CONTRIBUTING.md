# Contributing

## Pull requests only

**DON'T** push to the master branch directly. Always use pull requests and
let people discuss changes in pull request.

Pull requests should only be merged after all discussions have been
concluded and at least 2 reviewers gave their approvals.

## Changelog for major changes

When your pull request does major changes, please also add an entry to the changelog.

## Branch Naming Convention

We have defined some naming conventions for branches, according to their purpose. To guarantee the quality of the development process, we created a git hook to validate the names.

The [pre-commit](.githooks/pre-commit) hook will check if your branch name contains one of the following words:
* feature
* fix

We decided not to use *release branches* here, because the Solution Delivery team does not have long development cycles in their projects.

Basically, you must follow the following specification:

![Creating a new branch](docs/regex_branch_name.png)

### Examples

- *Good*
  - feature/ST-15-new-export-writer-xml
  - fix/ST-1509-download-big-files
  - fix/ST-50-upload-csv-with-break-lines


- *Bad*

  - update-project
  - teamapps-6065
  - fix
  - feature/ST-12

## Continuous Integration (CI/CD)

Aiming at a standardized and automated software delivery process, we developed a Continuous Integration process divided by steps.

* `build`  : build the app binary
* `lint`   : run the lint verification
* `audit`  : look for and report dependencies vulnerabilities
* `test`   : run tests and coverage reports
* `package`: build the app container image
* `push`   : tag and push the image to the registry
* `pages`  : publish coverage reports and swagger documentation to gitlab pages

The steps `build`, `lint`, `audit`, `test` and `package` will run when committing to any branch (including master) and will not trigger when committing to tags that follow the [rules](#tags).

* The `push` step will only be triggered when a new git tag is created, representing a new version that must be published. The tag name must follow the defined [rules](#tags).
* The `pages` step will only be triggered when there is a commit on the master.


### Tags

The [CI configuration](.gitlab-ci.yml) has some specifications about how to name your tags.

Please follow these rules:

* When tagging a release candidate, use this format: `1.0.0-rc.1`
  * When tagging another candidate for the same release, increase the last digit: `1.0.0-rc.2`

* When tagging a new release, use this format: `1.0.0`

For more details, check [Semantic Versioning Specification](https://semver.org/).
