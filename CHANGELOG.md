# [0.1.0](https://github.com/capsa-gg/capsa/compare/v0.0.7...v0.1.0) (2025-02-16)

## [0.0.7](https://github.com/capsa-gg/capsa/compare/v0.0.6...v0.0.7) (2025-02-16)


### Bug Fixes

* **deployment:** correctly update Chart versions ([310d348](https://github.com/capsa-gg/capsa/commit/310d3486feea503ddcd96270fb1e5c1296f1e03f))
* **deployment:** correctly update Chart versions ([#242](https://github.com/capsa-gg/capsa/issues/242)) ([d411579](https://github.com/capsa-gg/capsa/commit/d41157909c8c3995b42dd55b45523cc93c5440f9))
* **server:** capture logs as Sentry breadcrumbs if level below warning ([a2bdafd](https://github.com/capsa-gg/capsa/commit/a2bdafd499995f77b01944d3bf5c7b9a57f7b245))
* **web:** include first and last line for each log in merged response ([87fef32](https://github.com/capsa-gg/capsa/commit/87fef329dfd39940ce0c4db409c85eacd6ac59bc))


### Features

* opt-in support for Sentry ([a610c4f](https://github.com/capsa-gg/capsa/commit/a610c4ff4d1f3331626563c43b5d960b8e0a8bfe))
* opt-in support for Sentry ([#170](https://github.com/capsa-gg/capsa/issues/170)) ([a9979cf](https://github.com/capsa-gg/capsa/commit/a9979cff7c51ce3940fc5e04af0da4acbcd437b3))
* support merged log streaming ([d841747](https://github.com/capsa-gg/capsa/commit/d8417472a53511ad5fb7c9f6404aaf625d9b6562))

## [0.0.6](https://github.com/capsa-gg/capsa/compare/v0.0.5...v0.0.6) (2024-12-11)


### Bug Fixes

* **server:** don't allow deactivated users to be redeactivated ([aabc561](https://github.com/capsa-gg/capsa/commit/aabc561e7c1e137eb97339da91255134fa634595))
* **server:** honour chunk filtering without severity setting ([6d3cbed](https://github.com/capsa-gg/capsa/commit/6d3cbed89a297e97138ab3ce97a0ae099a725e0c))
* **web:** correctly redirect after login ([b152e89](https://github.com/capsa-gg/capsa/commit/b152e89b428b21402f60120dc9aea3c5f66841f3))
* **web:** login background not white anymore on login screen ([fafa0f6](https://github.com/capsa-gg/capsa/commit/fafa0f688c03c007178402c6676393d078f218b6))


### Features

* add log filtering by platform and log type ([3e06fdf](https://github.com/capsa-gg/capsa/commit/3e06fdf9a798c4a649cb538a808d75d19723eb85))
* add support for Game and Editor log types ([e86014b](https://github.com/capsa-gg/capsa/commit/e86014b626ad2e581ccc9cdeacd2f9d2416ecc56))
* **deployment:** add GHCR repos for DO Apps Spec ([895d4e9](https://github.com/capsa-gg/capsa/commit/895d4e9485b8b3ea8ce16c5d3776b7495d586faf))
* **deployment:** automatic db migration with server App Service spec ([dfa2602](https://github.com/capsa-gg/capsa/commit/dfa26021342fab39fced4592d3d73978f52918a4))
* **deployment:** automatic db migration with server App Service spec ([#122](https://github.com/capsa-gg/capsa/issues/122)) ([9e1a889](https://github.com/capsa-gg/capsa/commit/9e1a889bf3c3fc4457418163cc74e4774563c6c0))
* **server:** add endpoints for title and environment management ([599eb62](https://github.com/capsa-gg/capsa/commit/599eb625cde4609c72ef8cab4abe2c202485ebc0))
* **server:** add filtering for log overview ([8dad79b](https://github.com/capsa-gg/capsa/commit/8dad79b456999942fd5e2093b84b770eb9bafa6a))
* **server:** add user management endpoints ([7503aa4](https://github.com/capsa-gg/capsa/commit/7503aa4acec09119e70b4704b4bb8cbb4194158b))
* **server:** add user roles and deactivation/reactivation ([9b78410](https://github.com/capsa-gg/capsa/commit/9b784103446a7053cd4e7c1e6ca7ff6915f774a8))
* **server:** add zlib compression support ([0e65929](https://github.com/capsa-gg/capsa/commit/0e65929dca330e5210e93e7cd9fad1c44da50a3d))
* **server:** add zlib compression support ([#151](https://github.com/capsa-gg/capsa/issues/151)) ([5650cdb](https://github.com/capsa-gg/capsa/commit/5650cdbdc5069ad511c9e78a0aff3810194735b3))
* **server:** nightly automatic removal of old logs ([c06b80c](https://github.com/capsa-gg/capsa/commit/c06b80ca6184804fc5608038517a147c2d7c5a66))
* **server:** nightly automatic removal of old logs ([#123](https://github.com/capsa-gg/capsa/issues/123)) ([255167f](https://github.com/capsa-gg/capsa/commit/255167fd82748ce0bad34fdfeae1415926c479d7)), closes [#74](https://github.com/capsa-gg/capsa/issues/74)
* **server:** search endpoint ([2f3ad80](https://github.com/capsa-gg/capsa/commit/2f3ad80a4f8434e153c317f2f47e52b63156fdda))
* show linked log count for logs ([24a1126](https://github.com/capsa-gg/capsa/commit/24a112631ac46f29617230d108220e3cbc706d82))
* show linked log count for logs ([#149](https://github.com/capsa-gg/capsa/issues/149)) ([bcc7ced](https://github.com/capsa-gg/capsa/commit/bcc7cede363d78db037ddc8cfa11806373a64bf5))
* **web:** add admin pages and improve sidebar ([5a51dd2](https://github.com/capsa-gg/capsa/commit/5a51dd2606edeb6cc5eb969721abcdc4b886c7d5))
* **web:** add environment filtering for log overview ([010b99e](https://github.com/capsa-gg/capsa/commit/010b99e08bcb21d4f3f66d67d58c01b84ea7d5c5))
* **web:** add error boundary in web app ([7dff9dc](https://github.com/capsa-gg/capsa/commit/7dff9dc5b2aec63e2b510645d12ec5678b60efdc))
* **web:** add single log metadata tab ([0d2817c](https://github.com/capsa-gg/capsa/commit/0d2817c837c893c8433ce22f45f21f3b37efe40d))
* **web:** add snackbar notifications (by extending error snackbar) ([aef6a74](https://github.com/capsa-gg/capsa/commit/aef6a74361e7b7db05894add69320c9e69269ec6))
* **web:** dark mode support ([8da4bab](https://github.com/capsa-gg/capsa/commit/8da4bab4f17f3d56b4acdc6c901b2654d76eb983))
* **web:** search functionality ([6544414](https://github.com/capsa-gg/capsa/commit/6544414ea250285752a441efe2fc1baed1aafc66))
* **web:** title and environment management UI ([1d6f4a2](https://github.com/capsa-gg/capsa/commit/1d6f4a207d4f07368951e4e88e68bb23eb096259))
* **web:** user management page implementation ([4a0c17e](https://github.com/capsa-gg/capsa/commit/4a0c17e897dafc5cb09bc2ba371ddf98f6fa9d91))

## [0.0.5](https://github.com/capsa-gg/capsa/compare/v0.0.4...v0.0.5) (2024-11-20)


### Bug Fixes

* **server:** allow clients to access headers ([1bceed3](https://github.com/capsa-gg/capsa/commit/1bceed38c25925f20a6fd168a23a651682d1f58f))
* **server:** include number of unprocessed lines in chunk line count ([47d2edd](https://github.com/capsa-gg/capsa/commit/47d2edd5e4705bc0354aea25c285c6d01b270bd9))
* **web:** show all times in UTC ([f51152d](https://github.com/capsa-gg/capsa/commit/f51152da81cc4581def18c000e66ed784537f4ff))


### Features

* **server:** add absolute line numbers for filtered logs ([2b118c1](https://github.com/capsa-gg/capsa/commit/2b118c15ccb9ea54da0f3e1a772f12b85086a9d6))
* **server:** add basic log data to metadata api endpoint ([4ce95a0](https://github.com/capsa-gg/capsa/commit/4ce95a0959e96abcf028ff5070ad56b48ad9e542))
* **server:** support gzip log uploads ([d2680de](https://github.com/capsa-gg/capsa/commit/d2680de621774e53b4f9eee2f17ee19772d340ec))
* **server:** support log lines in log streaming with url params ([f9cb388](https://github.com/capsa-gg/capsa/commit/f9cb388e55c726cb3a16d093c1ac03c0b9c20389))
* show chunk count with log metadata ([01a90d7](https://github.com/capsa-gg/capsa/commit/01a90d7e00274c4596154d45399283947e6bc731))
* **web:** add absolute line numbers for filtered logs ([aa4ee4b](https://github.com/capsa-gg/capsa/commit/aa4ee4b888ef034146c05a344085e86f881b4f64))
* **web:** add copy url button for sharing log view with filters ([c6922f8](https://github.com/capsa-gg/capsa/commit/c6922f8e627d4d42ccfa780483a2459d91b21a26))
* **web:** add filters to URL bar on single log ([3b91bd3](https://github.com/capsa-gg/capsa/commit/3b91bd31958c4cdfd735fdb04e45ffad4c911168))
* **web:** add single log line filtering by severity ([c2a5f4e](https://github.com/capsa-gg/capsa/commit/c2a5f4e77b83a114969b52174c3932f09a839c05))
* **web:** display log metadata on single log page ([8f8276d](https://github.com/capsa-gg/capsa/commit/8f8276d1f29fb5c593835e49d0a4b7ca637234be))
* **web:** filter log lines by included and excluded categories ([3b3d13d](https://github.com/capsa-gg/capsa/commit/3b3d13d44576d27ebc5d32ea1dad375c5b1ffa69))
* **web:** use a Worker for fetching and processing log data ([bfba4c0](https://github.com/capsa-gg/capsa/commit/bfba4c04297dbf3b7784403a7251132abf2ef690))


### Performance Improvements

* **server:** do not fetch chunks that can be ignored ([310b356](https://github.com/capsa-gg/capsa/commit/310b356de70eff330c417f9e314e6ae09516adad))

## [0.0.4](https://github.com/capsa-gg/capsa/compare/v0.0.3...v0.0.4) (2024-11-18)


### Bug Fixes

* **server:** correctly reset log fields on each request for auth middleware ([412e85e](https://github.com/capsa-gg/capsa/commit/412e85eded1d2a5000755bf14b6f1fecf1a02f50))
* **web:** correct header names in log list ([70a18e1](https://github.com/capsa-gg/capsa/commit/70a18e1ca9311dd6c04382dc33f554adc91a3d9f))


### chore

* rename project ([7a604c5](https://github.com/capsa-gg/capsa/commit/7a604c5f0b9d9ec87083212f4d69603b327c6265))


### Features

* **ci:** add example Helm charts ([96bcadc](https://github.com/capsa-gg/capsa/commit/96bcadc45a60d3c0541edeb679db3037d1252665))
* **deployment:** add DigitalOcean Apps specs ([11e8ab3](https://github.com/capsa-gg/capsa/commit/11e8ab31c9f5bbc168ab95804458fb11356b6292))
* **server:** add log chunk processing for incoming logs ([2555702](https://github.com/capsa-gg/capsa/commit/2555702ffac4a73eb168c199a56eb7508d34a2d0))
* **server:** add log chunk processing logic with unit tests ([5cd088d](https://github.com/capsa-gg/capsa/commit/5cd088dbff9162d226fe5668cd13bf210db8f4b5))
* **server:** add version command ([4a945e6](https://github.com/capsa-gg/capsa/commit/4a945e631cddede188469c95bd3a7b4ad14bfaff))
* **server:** expose log data in overview api ([cc7d71b](https://github.com/capsa-gg/capsa/commit/cc7d71bc8e15978bd49362733251076f9bf77b91))
* **server:** include log link for clients on log session creation ([fa2c6b2](https://github.com/capsa-gg/capsa/commit/fa2c6b2a8227dddd8036bee23924bb5051fb9e7a))
* **server:** serve static files and include logo in emails ([53422d8](https://github.com/capsa-gg/capsa/commit/53422d895ca2d2cf75b8a2cff1b36a63e4a781b7))
* **web:** add (temp) branding ([e0d21f7](https://github.com/capsa-gg/capsa/commit/e0d21f7796d8f7ab842bf8a907ce96667bae7e4d))
* **web:** add custom colors to Monaco editor ([77005b1](https://github.com/capsa-gg/capsa/commit/77005b181b156e64d43dbf11e69ad4fd1153f17c))
* **web:** add Monaco editor for viewing logs ([4217345](https://github.com/capsa-gg/capsa/commit/42173452568b417e42b396af7906e1bdf2b0d4b3))
* **web:** show log info in overview ([dc188d4](https://github.com/capsa-gg/capsa/commit/dc188d4706e05eeea04b3af7571d528297348a34))
* **web:** support redirecting after login ([35e37a2](https://github.com/capsa-gg/capsa/commit/35e37a26a63c5e66048a38bd2a55800f82c115a1))


### Performance Improvements

* **server:** add benchmarking code for chunk processing of 1M lines ([b96ad72](https://github.com/capsa-gg/capsa/commit/b96ad72cbe55c7497b15889eb5e0d58e1a0db78a))


### BREAKING CHANGES

* moved project

Signed-off-by: Luciano Nooijen <luciano@lucianonooijen.com>

## [0.0.3](https://github.com/capsa-gg/capsa/compare/v0.0.2...v0.0.3) (2024-11-11)


### Bug Fixes

* better hot reloading ([adfad39](https://github.com/capsa-gg/capsa/commit/adfad39960d42fb39cefda7c056a77ddd30731a3))
* build correct Docker image on release ([3c39aff](https://github.com/capsa-gg/capsa/commit/3c39aff67072e31e11fb3e5ed31003084b9a4574))
* log insertion improvements ([5db65c0](https://github.com/capsa-gg/capsa/commit/5db65c0d79d1f2a5b77a1defe2d69e25d96afd5c))


### Features

* **server:** add security headers middleware ([a7eaaff](https://github.com/capsa-gg/capsa/commit/a7eaaff25d8b52d023557e2e1656a5be89ac9108))
* **web:** add security headers in Next config ([08823e5](https://github.com/capsa-gg/capsa/commit/08823e5b223046578da111af39c0ca1db8f9405f))
* **webapp:** dynamic loading of server url instead of inlining NEXT_PUBLIC_ variables into html ([60a75dd](https://github.com/capsa-gg/capsa/commit/60a75dde7b9446a09920be3181619b7cfd5d8e7a))
* **web:** display environment info on homepage ([2746d1f](https://github.com/capsa-gg/capsa/commit/2746d1f7fd7e2911003ed7e55044a422ec0133c4))
* **web:** log overview and single log pages ([52eb468](https://github.com/capsa-gg/capsa/commit/52eb46802edbe26bcbd38f1f6bae0f7896405b89))

## [0.0.2](https://github.com/capsa-gg/capsa/compare/v0.0.1...v0.0.2) (2024-11-10)


### Bug Fixes

* add assets to release ([75ba91f](https://github.com/capsa-gg/capsa/commit/75ba91fed15e6d2960d117dee5d1f06775ca9a68))
* deployed swagger trailing slash removed ([80e1113](https://github.com/capsa-gg/capsa/commit/80e111328a2962b998019b2f98e71b786963666b))
* remove port on dev mode that is not localhost ([7847b8d](https://github.com/capsa-gg/capsa/commit/7847b8de815d58c45cad3f567d996022c53749fe))
* remove unused dependencies ([70dd4f8](https://github.com/capsa-gg/capsa/commit/70dd4f8c74cff5890f0eb398b0a116d1061ec112))
* semantic release ([f71b529](https://github.com/capsa-gg/capsa/commit/f71b529fc8772d672c4451ae74a84739dfe28912))
* **server:** correct swagger docs ([b66d629](https://github.com/capsa-gg/capsa/commit/b66d629ebcad0693dc3f031231f8a1402b82c39e))
* **server:** fix failing ci ([6fb98bc](https://github.com/capsa-gg/capsa/commit/6fb98bcfff7ee0c23cae4d16c6e5d92366e56c74))
* **server:** formatting ([d88ec51](https://github.com/capsa-gg/capsa/commit/d88ec51520121475fb81e622e269429162b3e6c3))
* **server:** set correct endpoint for Swagger in dev mode ([f259347](https://github.com/capsa-gg/capsa/commit/f259347c1b849fd3d8c8179927bad10f9b1e22d7))
* set correct job names, add todos ([517529d](https://github.com/capsa-gg/capsa/commit/517529ded19a05adaedbd4a4406f03b82f76f933))
* swagger links ([1da1f59](https://github.com/capsa-gg/capsa/commit/1da1f59a8693349d7c77757bf5ebaa61ca260d8c))
* **webapp:** temporary workaround to get webapp working ([442b825](https://github.com/capsa-gg/capsa/commit/442b8259cc999509d9d49623b44d47f1cd95ac6c))


### Features

* **server:** add Gin server with Swagger docs ([9b86f79](https://github.com/capsa-gg/capsa/commit/9b86f793a95281a7c64c412708d021fdb753ac38))
* **server:** cli command for listing all environments ([d073e85](https://github.com/capsa-gg/capsa/commit/d073e8538129360d63d08ee8c457fe4b30c5d410))
* **server:** cli command for listing all environments ([c806656](https://github.com/capsa-gg/capsa/commit/c80665691656869a9a2181471a71a91f695343f2))
* **server:** cli command for listing all environments ([7b82255](https://github.com/capsa-gg/capsa/commit/7b82255bf37fca4c5b3ef2e64769d6aff6584819))
* **server:** client auth endpoint ([97f1ce2](https://github.com/capsa-gg/capsa/commit/97f1ce24b83b051541bf35c6a19b979aa88daf21))
* **server:** domain logic for authentication and adding log sessions ([858a497](https://github.com/capsa-gg/capsa/commit/858a49739f1559d947148bf2970cb2e1e512158d))
* **server:** generating jwks, add .well-known/jwks.json endpoint ([c303758](https://github.com/capsa-gg/capsa/commit/c3037581c04b6052ecaa19188274c228716408f4))
* **server:** get log metadata from database ([a194346](https://github.com/capsa-gg/capsa/commit/a194346a39c077f15dbe71de6a0538d5da7f59f6))
* **server:** hot reloading in development ([1eb94fb](https://github.com/capsa-gg/capsa/commit/1eb94fb7cc3fb37ccdfb573c6d15d02d490b96bb))
* **server:** include server version in requests ([1cd22cb](https://github.com/capsa-gg/capsa/commit/1cd22cb2001503ffae1dc4a5d80e984e24fc67ed))
* **server:** jwt generation ([07c19ea](https://github.com/capsa-gg/capsa/commit/07c19eacb6ff90d73de089238797ea9bcf1c67ef))
* **server:** log and environment listing ([2d61952](https://github.com/capsa-gg/capsa/commit/2d619524b52d4af38af418cec8ba4ccfcaca2631))
* **server:** log metadata saving ([091ecad](https://github.com/capsa-gg/capsa/commit/091ecad013abf59840fb8e57427f490fc928ca8c))
* **server:** password reset flow ([e22e15c](https://github.com/capsa-gg/capsa/commit/e22e15c15fe6f6b98d697f7cc078dea2e83a2e0e))
* **server:** store log chunks without processing ([85bcb81](https://github.com/capsa-gg/capsa/commit/85bcb81162e6e814afd80aabc5c572059c88e0f9))
* **server:** stream log chunks for reading ([8c3cc2f](https://github.com/capsa-gg/capsa/commit/8c3cc2f7e82bbb665ab6e5e57439b09958a8d04f))
* **server:** transactional emails ([4b813d1](https://github.com/capsa-gg/capsa/commit/4b813d1ed9178e76a3f8350e2f151b08b33478b7))
* **server:** user login route ([8e0cd23](https://github.com/capsa-gg/capsa/commit/8e0cd237b6ed3f8e54c87c3d343363e0eeb9cceb))
* **webapp:** auth pages, middleware, api call hooks ([9c8eeaa](https://github.com/capsa-gg/capsa/commit/9c8eeaa60c05de25f09ee04bbec1dde29e0d73cb))
* **webapp:** auth pages, middleware, api call hooks ([9d6e942](https://github.com/capsa-gg/capsa/commit/9d6e94252071cbdbffacdc07b47ce6a51eb6431a))

## 0.0.1 (2024-11-10)

Project skeleton setup
