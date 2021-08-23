PingAccess SDK Go
==================

- [![Gitter](https://badges.gitter.im/iwarapter/pingaccess-sdk-go.svg)](https://gitter.im/iwarapter/pingaccess-sdk-go?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
  [![Coverage](https://sonarcloud.io/api/project_badges/measure?project=github.com.iwarapter.pingaccess-sdk-go&metric=coverage)](https://sonarcloud.io/dashboard?id=github.com.iwarapter.pingaccess-sdk-go)
  [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=github.com.iwarapter.pingaccess-sdk-go&metric=alert_status)](https://sonarcloud.io/dashboard?id=github.com.iwarapter.pingaccess-sdk-go)
  [![Build Status](https://travis-ci.org/iwarapter/pingaccess-sdk-go.svg?branch=master)](https://travis-ci.org/iwarapter/pingaccess-sdk-go)

This project was created to support the [terraform-provider-pingaccess](https://github.com/iwarapter/terraform-provider-pingaccess) and other experimental projects.

The SDK is currently generated, however the PingAccess API docs have several mistakes which are handled and documented below

The PingAccess Admin API introduces breaking changes between minor product updates, to handle this each major/minor version is published as its own module version.

For example `github.com/iwarapter/pingaccess-sdk-go/v60@6.0` is PingAccess version 6.0 and is tracked on the 6.0 branch.

Using the SDK
----------------------

```
go get github.com/iwarapter/pingaccess-sdk-go/v60@6.0
```

Testing the SDK
---------------------------

In order to test the provider, you can run `make test`.

```sh
$ make test
```

Missing Models
--------------
PolicyItem, this was removed from the documentation between 6.0.1 and 6.0.2.0 this is the reference object for rules on the `ApplicationView` and `ResourceView`.
```go
type PolicyItem struct {
	Id   json.Number `json:"id,omitempty"`
	Type *string     `json:"type,omitempty"`
}
```

Part of the KeyPair GetKeypairsCreatableGeneralNamesCommand response, this is also missing from the API docs.
```go
type SanType struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}
```

Part of the Config Export API used by the JsonWebKey, haven't tested this so empty struct for now.
```
type PublicKey struct {
}
```

Missing Model Fields
--------------------

- `AcmeServerView` has undefined field `id` which has been assigned the type `string`
- `AccessTokenValidatorView` has undefined field `id` which has been assigned the type `json.Number`
- `AcmeAccountView` has undefined field `id` which has been assigned the type `string`
- `AcmeCertificateRequestView` has undefined field `id` which has been assigned the type `string`
- `AgentView` has undefined field `id` which has been assigned the type `json.Number`
- `ApplicationView` has undefined field `id` which has been assigned the type `json.Number`
- `AuthnReqListView` has undefined field `id` which has been assigned the type `json.Number`
- `AvailabilityProfileView` has undefined field `id` which has been assigned the type `json.Number`
- `EngineListenerView` has undefined field `id` which has been assigned the type `json.Number`
- `EngineView` has undefined field `id` which has been assigned the type `json.Number`
- `ExportData` has undefined field `id` which has been assigned the type `json.Number`
- `GlobalUnprotectedResourceView` has undefined field `id` which has been assigned the type `json.Number`
- `HsmProviderView` has undefined field `id` which has been assigned the type `json.Number`
- `HttpClientProxyView` has undefined field `id` which has been assigned the type `json.Number`
- `HttpsListenerView` has undefined field `id` which has been assigned the type `json.Number`
- `IdentityMappingView` has undefined field `id` which has been assigned the type `json.Number`
- `LoadBalancingStrategyView` has undefined field `id` which has been assigned the type `json.Number`
- `RejectionHandlerView` has undefined field `id` which has been assigned the type `json.Number`
- `ReplicaAdminView` has undefined field `id` which has been assigned the type `json.Number`
- `ResourceOrderView` has undefined field `id` which has been assigned the type `json.Number`
- `ResourceView` has undefined field `id` which has been assigned the type `json.Number`
- `RuleSetView` has undefined field `id` which has been assigned the type `json.Number`
- `RuleView` has undefined field `id` which has been assigned the type `json.Number`
- `SharedSecretView` has undefined field `id` which has been assigned the type `json.Number`
- `SiteAuthenticatorView` has undefined field `id` which has been assigned the type `json.Number`
- `SiteView` has undefined field `id` which has been assigned the type `json.Number`
- `ThirdPartyServiceView` has undefined field `id` which has been assigned the type `string`
- `TrustedCertificateGroupView` has undefined field `id` which has been assigned the type `json.Number`
- `UserPasswordView` has undefined field `id` which has been assigned the type `json.Number`
- `UserView` has undefined field `id` which has been assigned the type `json.Number`
- `VirtualHostView` has undefined field `id` which has been assigned the type `json.Number`
- `WebSessionView` has undefined field `id` which has been assigned the type `json.Number`

Many of the API's do not document the ID field for each of the API's and use several different types:
- `int` literal
- `string` numeric literal (`"1"`)
- `string` uuid


Incorrent Model Fields
----------------------

- `ChainCertificateView`
- - `expires` Type: `int` not `string`
- - `validFrom` Type: `int` not `string`
- `KeyPairView`
- - `expires` Type: `int` not `string`
- - `validFrom` Type: `int` not `string`
- `TrustedCertView`
- - `expires` Type: `int` not `string`
- - `validFrom` Type: `int` not `string`
