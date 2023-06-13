"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[6088],{3905:(e,t,n)=>{n.d(t,{Zo:()=>m,kt:()=>u});var s=n(67294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);t&&(s=s.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,s)}return n}function r(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,s,o=function(e,t){if(null==e)return{};var n,s,o={},a=Object.keys(e);for(s=0;s<a.length;s++)n=a[s],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(s=0;s<a.length;s++)n=a[s],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var l=s.createContext({}),p=function(e){var t=s.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):r(r({},t),e)),n},m=function(e){var t=p(e.components);return s.createElement(l.Provider,{value:t},e.children)},c={inlineCode:"code",wrapper:function(e){var t=e.children;return s.createElement(s.Fragment,{},t)}},d=s.forwardRef((function(e,t){var n=e.components,o=e.mdxType,a=e.originalType,l=e.parentName,m=i(e,["components","mdxType","originalType","parentName"]),d=p(n),u=o,h=d["".concat(l,".").concat(u)]||d[u]||c[u]||a;return n?s.createElement(h,r(r({ref:t},m),{},{components:n})):s.createElement(h,r({ref:t},m))}));function u(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=n.length,r=new Array(a);r[0]=d;var i={};for(var l in t)hasOwnProperty.call(t,l)&&(i[l]=t[l]);i.originalType=e,i.mdxType="string"==typeof e?e:o,r[1]=i;for(var p=2;p<a;p++)r[p]=n[p];return s.createElement.apply(null,r)}return s.createElement.apply(null,n)}d.displayName="MDXCreateElement"},35032:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>r,default:()=>c,frontMatter:()=>a,metadata:()=>i,toc:()=>p});var s=n(87462),o=(n(67294),n(3905));const a={sidebar_position:1},r="Testing",i={unversionedId:"building-modules/testing",id:"building-modules/testing",title:"Testing",description:"The Cosmos SDK contains different types of tests.",source:"@site/docs/building-modules/16-testing.md",sourceDirName:"building-modules",slug:"/building-modules/testing",permalink:"/main/building-modules/testing",draft:!1,tags:[],version:"current",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"Modules depinject-ready",permalink:"/main/building-modules/depinject"},next:{title:"Overview of app.go",permalink:"/main/building-apps/app-go"}},l={},p=[{value:"Unit Tests",id:"unit-tests",level:2},{value:"Example",id:"example",level:3},{value:"Integration Tests",id:"integration-tests",level:2},{value:"Example",id:"example-1",level:3},{value:"Deterministic and Regression tests",id:"deterministic-and-regression-tests",level:2},{value:"Simulations",id:"simulations",level:2},{value:"End-to-end Tests",id:"end-to-end-tests",level:2},{value:"Learn More",id:"learn-more",level:2}],m={toc:p};function c(e){let{components:t,...n}=e;return(0,o.kt)("wrapper",(0,s.Z)({},m,n,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"testing"},"Testing"),(0,o.kt)("p",null,"The Cosmos SDK contains different types of ",(0,o.kt)("a",{parentName:"p",href:"https://martinfowler.com/articles/practical-test-pyramid.html"},"tests"),".\nThese tests have different goals and are used at different stages of the development cycle.\nWe advice, as a general rule, to use tests at all stages of the development cycle.\nIt is adviced, as a chain developer, to test your application and modules in a similar way than the SDK."),(0,o.kt)("p",null,"The rationale behind testing can be found in ",(0,o.kt)("a",{parentName:"p",href:"https://docs.cosmos.network/main/architecture/adr-059-test-scopes.html"},"ADR-59"),"."),(0,o.kt)("h2",{id:"unit-tests"},"Unit Tests"),(0,o.kt)("p",null,"Unit tests are the lowest test category of the ",(0,o.kt)("a",{parentName:"p",href:"https://martinfowler.com/articles/practical-test-pyramid.html"},"test pyramid"),".\nAll packages and modules should have unit test coverage. Modules should have their dependencies mocked: this means mocking keepers."),(0,o.kt)("p",null,"The SDK uses ",(0,o.kt)("inlineCode",{parentName:"p"},"mockgen")," to generate mocks for keepers:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/scripts/mockgen.sh#L3-L6\n")),(0,o.kt)("p",null,"You can read more about mockgen ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/golang/mock"},"here"),"."),(0,o.kt)("h3",{id:"example"},"Example"),(0,o.kt)("p",null,"As an example, we will walkthrough the ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/gov/keeper/keeper_test.go"},"keeper tests")," of the ",(0,o.kt)("inlineCode",{parentName:"p"},"x/gov")," module."),(0,o.kt)("p",null,"The ",(0,o.kt)("inlineCode",{parentName:"p"},"x/gov")," module has a ",(0,o.kt)("inlineCode",{parentName:"p"},"Keeper")," type requires a few external dependencies (ie. imports outside ",(0,o.kt)("inlineCode",{parentName:"p"},"x/gov")," to work properly)."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/gov/keeper/keeper.go#L61-L65\n")),(0,o.kt)("p",null,"In order to only test ",(0,o.kt)("inlineCode",{parentName:"p"},"x/gov"),", we mock the ",(0,o.kt)("a",{parentName:"p",href:"https://docs.cosmos.network/v0.46/building-modules/keeper.html#type-definition"},"expected keepers")," and instantiate the ",(0,o.kt)("inlineCode",{parentName:"p"},"Keeper")," with the mocked dependencies. Note that we may need to configure the mocked dependencies to return the expected values:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/gov/keeper/common_test.go#L67-L81\n")),(0,o.kt)("p",null,"This allows us to test the ",(0,o.kt)("inlineCode",{parentName:"p"},"x/gov")," module without having to import other modules."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/gov/keeper/keeper_test.go#L3-L35\n")),(0,o.kt)("p",null,"We can test then create unit tests using the newly created ",(0,o.kt)("inlineCode",{parentName:"p"},"Keeper")," instance."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/gov/keeper/keeper_test.go#L73-L91\n")),(0,o.kt)("h2",{id:"integration-tests"},"Integration Tests"),(0,o.kt)("p",null,"Integration tests are at the second level of the ",(0,o.kt)("a",{parentName:"p",href:"https://martinfowler.com/articles/practical-test-pyramid.html"},"test pyramid"),".\nIn the SDK, we locate our integration tests under ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/tree/main/tests/integration"},(0,o.kt)("inlineCode",{parentName:"a"},"/tests/integrations")),"."),(0,o.kt)("p",null,"The goal of these integration tests is to test how a component interacts with other dependencies. Compared to unit tests, integration tests do not mock dependencies. Instead, they use the direct dependencies of the component. This differs as well from end-to-end tests, which test the component with a full application."),(0,o.kt)("p",null,"Integration tests interact with the tested module via the defined ",(0,o.kt)("inlineCode",{parentName:"p"},"Msg")," and ",(0,o.kt)("inlineCode",{parentName:"p"},"Query")," services. The result of the test can be verified by checking the state of the application, by checking the emitted events or the response. It is adviced to combine two of these methods to verify the result of the test."),(0,o.kt)("p",null,"The SDK provides small helpers for quickly setting up an integration tests. These helpers can be found at ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/blob/main/testutil/integration"},"https://github.com/cosmos/cosmos-sdk/blob/main/testutil/integration"),"."),(0,o.kt)("h3",{id:"example-1"},"Example"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/29e22b3bdb05353555c8e0b269311bbff7b8deca/testutil/integration/example_test.go#L22-L89\n")),(0,o.kt)("h2",{id:"deterministic-and-regression-tests"},"Deterministic and Regression tests"),(0,o.kt)("p",null,"Tests are written for queries in the Cosmos SDK which have ",(0,o.kt)("inlineCode",{parentName:"p"},"module_query_safe")," Protobuf annotation."),(0,o.kt)("p",null,"Each query is tested using 2 methods:"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"Use property-based testing with the ",(0,o.kt)("a",{parentName:"li",href:"https://pkg.go.dev/pgregory.net/rapid@v0.5.3"},(0,o.kt)("inlineCode",{parentName:"a"},"rapid"))," library. The property that is tested is that the query response and gas consumption are the same upon 1000 query calls."),(0,o.kt)("li",{parentName:"ul"},"Regression tests are written with hardcoded responses and gas, and verify they don't change upon 1000 calls and between SDK patch versions.")),(0,o.kt)("p",null,"Here's an example of regression tests:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/tests/integration/bank/keeper/deterministic_test.go#L102-L115\n")),(0,o.kt)("h2",{id:"simulations"},"Simulations"),(0,o.kt)("p",null,"Simulations uses as well a minimal application, built with ",(0,o.kt)("a",{parentName:"p",href:"/main/packages/depinject"},(0,o.kt)("inlineCode",{parentName:"a"},"depinject")),":"),(0,o.kt)("admonition",{type:"note"},(0,o.kt)("p",{parentName:"admonition"},"You can as well use the ",(0,o.kt)("inlineCode",{parentName:"p"},"AppConfig")," ",(0,o.kt)("inlineCode",{parentName:"p"},"configurator")," for creating an ",(0,o.kt)("inlineCode",{parentName:"p"},"AppConfig")," ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/slashing/app_test.go#L54-L62"},"inline"),". There is no difference between those two ways, use whichever you prefer.")),(0,o.kt)("p",null,"Following is an example for ",(0,o.kt)("inlineCode",{parentName:"p"},"x/gov/")," simulations:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/gov/simulation/operations_test.go#L292-L310\n")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-rc1/x/gov/simulation/operations_test.go#L69-L111\n")),(0,o.kt)("h2",{id:"end-to-end-tests"},"End-to-end Tests"),(0,o.kt)("p",null,"End-to-end tests are at the top of the ",(0,o.kt)("a",{parentName:"p",href:"https://martinfowler.com/articles/practical-test-pyramid.html"},"test pyramid"),".\nThey must test the whole application flow, from the user perspective (for instance, CLI tests). They are located under ",(0,o.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/tree/main/tests/e2e"},(0,o.kt)("inlineCode",{parentName:"a"},"/tests/e2e")),"."),(0,o.kt)("p",null,"For that, the SDK is using ",(0,o.kt)("inlineCode",{parentName:"p"},"simapp")," but you should use your own application (",(0,o.kt)("inlineCode",{parentName:"p"},"appd"),").\nHere are some examples:"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"SDK E2E tests: ",(0,o.kt)("a",{parentName:"li",href:"https://github.com/cosmos/cosmos-sdk/tree/main/tests/e2e"},"https://github.com/cosmos/cosmos-sdk/tree/main/tests/e2e"),"."),(0,o.kt)("li",{parentName:"ul"},"Cosmos Hub E2E tests: ",(0,o.kt)("a",{parentName:"li",href:"https://github.com/cosmos/gaia/tree/main/tests/e2e"},"https://github.com/cosmos/gaia/tree/main/tests/e2e"),"."),(0,o.kt)("li",{parentName:"ul"},"Osmosis E2E tests: ",(0,o.kt)("a",{parentName:"li",href:"https://github.com/osmosis-labs/osmosis/tree/main/tests/e2e"},"https://github.com/osmosis-labs/osmosis/tree/main/tests/e2e"),".")),(0,o.kt)("admonition",{title:"warning",type:"note"},(0,o.kt)("p",{parentName:"admonition"},"The SDK is in the process of creating its E2E tests, as defined in ",(0,o.kt)("a",{parentName:"p",href:"https://docs.cosmos.network/main/architecture/adr-059-test-scopes.html"},"ADR-59"),". This page will eventually be updated with better examples.")),(0,o.kt)("h2",{id:"learn-more"},"Learn More"),(0,o.kt)("p",null,"Learn more about testing scope in ",(0,o.kt)("a",{parentName:"p",href:"https://docs.cosmos.network/main/architecture/adr-059-test-scopes.html"},"ADR-59"),"."))}c.isMDXComponent=!0}}]);