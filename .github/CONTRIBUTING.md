# How to contribute

Contributions to Diru are highly encouraged and desired. Below are some guidelines that will help make the process as smooth as possible.

## Getting Started

- Make sure you have a [GitHub account](https://github.com/signup/free)
- Submit a new issue, assuming one does not already exist.
  - Clearly describe the issue including steps to reproduce when it is a bug.
  - Make sure you fill in the earliest version that you know has the issue.
- Fork the repository on GitHub

## Suggesting Enhancements

- When submitting an issue for an enhancement, please be as clear as possible about why you think the enhancement is needed and what the benefit of
    it would be.

## Making Changes

- From your fork of the repository, create a topic branch where work on your change will take place.
- To quickly create a topic branch based on main; `git checkout -b my_contribution main`. Please avoid working directly on the `main` branche.
- Make commits of logical units.
- Check for unnecessary whitespace with `git diff --check` before committing.
- Please follow the prevailing code conventions in the repository. Differences in style make the code harder to understand for everyone.
- Make sure you have added all the necessary unit tests for your changes.
- Run _all_ unit tests to assure nothing else was accidentally broken.
- Make sure your commit messages are in the proper format.
  - The first line of the commit message should be a short description around 50 characters in length and be prefixed with the GitHub issue it refers to with parentheses surrounding that. If the GitHub issue is #25, you should have `(#25)` prefixed to the message.
  - If the commit is about documentation, the message should be prefixed with `(doc)`.
  - If it is a trivial commit or one of formatting/spaces fixes, it should be prefixed with `(maint)`.
  - After the subject, skip one line and fill out a body if the subject line is not informative enough.
  - Sometimes you will find that even a tiny code change has a commit body that needs to be very detailed and make take more time to do than the actual change itself!
  - The body:
    - Should wrap at `72` characters.
    - Explains more fully the reason(s) for the change and contrasts with previous behavior.
    - Uses present tense. "Fix" versus "Fixed".
    - Contains a sign off statement to confirm that you agree to the [DCO](#DCO)

A good example of a commit message is as follows:

```text
(#123) Add more cowbell to Get-Something

The functionality of Get-Something would be greatly improved if there was a little
more 'pizzazz' added to it. I propose a cowbell. Adding more cowbell has been
shown in studies to both increase one's mojo, and cement one's status
as a rock legend.

Signed-off-by: Your Name <your@email.address>
```

## Documentation

We are infallible and as such the documentation needs no correction. In the highly
unlikely event that that is _not_ the case, commits to update or add documentation
are highly appreciated.

## Submitting Changes

- Push your changes to a topic branch in your fork of the repository.
- Submit a pull request to the main repository.
- Once the pull request has been reviewed and accepted, it will be merged with the master branch.
- Celebrate

## DCO
```text
Developer Certificate of Origin
Version 1.1

Copyright (C) 2004, 2006 The Linux Foundation and its contributors.
1 Letterman Drive
Suite D4700
San Francisco, CA, 94129

Everyone is permitted to copy and distribute verbatim copies of this
license document, but changing it is not allowed.


Developer's Certificate of Origin 1.1

By making a contribution to this project, I certify that:

(a) The contribution was created in whole or in part by me and I
    have the right to submit it under the open source license
    indicated in the file; or

(b) The contribution is based upon previous work that, to the best
    of my knowledge, is covered under an appropriate open source
    license and I have the right under that license to submit that
    work with modifications, whether created in whole or in part
    by me, under the same open source license (unless I am
    permitted to submit under a different license), as indicated
    in the file; or

(c) The contribution was provided directly to me by some other
    person who certified (a), (b) or (c) and I have not modified
    it.

(d) I understand and agree that this project and the contribution
    are public and that a record of the contribution (including all
    personal information I submit with it, including my sign-off) is
    maintained indefinitely and may be redistributed consistent with
    this project or the open source license(s) involved.
```
## Additional Resources

- [General GitHub documentation](https://help.github.com/)
- [GitHub forking documentation](https://guides.github.com/activities/forking/)
- [GitHub pull request documentation](https://help.github.com/send-pull-requests/)
- [GitHub Flow guide](https://guides.github.com/introduction/flow/)
- [GitHub's guide to contributing to open source projects](https://guides.github.com/activities/contributing-to-open-source/)
