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
