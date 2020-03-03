(window.webpackJsonp=window.webpackJsonp||[]).push([[45],{172:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return l})),r.d(t,"metadata",(function(){return o})),r.d(t,"rightToc",(function(){return p})),r.d(t,"default",(function(){return u}));var n=r(1),a=r(9),c=(r(0),r(181)),l={title:"Op Call [object]"},o={id:"opspec/reference/structure/op-directory/op/call/op",title:"Op Call [object]",description:"An object defining an op call.",source:"@site/docs/opspec/reference/structure/op-directory/op/call/op.md",permalink:"/docs/opspec/reference/structure/op-directory/op/call/op",editUrl:"https://github.com/opctl/opctl/edit/master/docs/docs/opspec/reference/structure/op-directory/op/call/op.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1583210148,sidebar:"docs",previous:{title:"Image [object]",permalink:"/docs/opspec/reference/structure/op-directory/op/call/container/image"},next:{title:"Parallel Loop Call [object]",permalink:"/docs/opspec/reference/structure/op-directory/op/call/parallel-loop"}},p=[{value:"Properties",id:"properties",children:[{value:"ref",id:"ref",children:[]},{value:"Example ref (opspec-pkgs golang.build.bin 2.0.0)",id:"example-ref-opspec-pkgs-golangbuildbin-200",children:[]},{value:"pullCreds",id:"pullcreds",children:[]},{value:"inputs",id:"inputs",children:[]},{value:"outputs",id:"outputs",children:[]}]}],b={rightToc:p},i="wrapper";function u(e){var t=e.components,r=Object(a.a)(e,["components"]);return Object(c.b)(i,Object(n.a)({},b,r,{components:t,mdxType:"MDXLayout"}),Object(c.b)("p",null,"An object defining an op call."),Object(c.b)("h2",{id:"properties"},"Properties"),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},"must have",Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"#ref"}),"ref")))),Object(c.b)("li",{parentName:"ul"},"may have",Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"#inputs"}),"inputs")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"#outputs"}),"outputs")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"#pullcreds"}),"pullCreds"))))),Object(c.b)("h3",{id:"ref"},"ref"),Object(c.b)("p",null,"A string referencing a local or remote operation."),Object(c.b)("p",null,"Must be one of:"),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},"an absolute path referencing an op embedded in the current op."),Object(c.b)("li",{parentName:"ul"},"a relative path referencing an op existing on the same local filesystem."),Object(c.b)("li",{parentName:"ul"},"a string in ",Object(c.b)("inlineCode",{parentName:"li"},"git-repo#{SEMVER_GIT_TAG}/path")," format referencing a network resolvable op.")),Object(c.b)("h3",{id:"example-ref-opspec-pkgs-golangbuildbin-200"},"Example ref (",Object(c.b)("a",Object(n.a)({parentName:"h3"},{href:"https://github.com/opspec-pkgs/golang.build.bin"}),"opspec-pkgs golang.build.bin 2.0.0"),")"),Object(c.b)("p",null,Object(c.b)("inlineCode",{parentName:"p"},"ref: 'github.com/opspec-pkgs/golang.build.bin#2.0.0'")),Object(c.b)("h3",{id:"pullcreds"},"pullCreds"),Object(c.b)("p",null,"An ",Object(c.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/opspec/reference/structure/op-directory/op/call/pull-creds"}),"pull-creds")," object defining creds used to pull the op from a private source."),Object(c.b)("h3",{id:"inputs"},"inputs"),Object(c.b)("p",null,"An object for which each key is an input of the referenced op and the value is one of:"),Object(c.b)("table",null,Object(c.b)("thead",{parentName:"table"},Object(c.b)("tr",{parentName:"thead"},Object(c.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"value"),Object(c.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"meaning"))),Object(c.b)("tbody",{parentName:"table"},Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"null"),Object(c.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Bind input to variable w/ same name (equivalent to ",Object(c.b)("inlineCode",{parentName:"td"},"$(INPUT_NAME)"),")")),Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(c.b)("a",Object(n.a)({parentName:"td"},{href:"/docs/opspec/reference/structure/op-directory/op/reference"}),"reference")),Object(c.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Bind referenced variable to the named input")),Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(c.b)("a",Object(n.a)({parentName:"td"},{href:"/docs/opspec/reference/structure/op-directory/op/initializer"}),"initializer")),Object(c.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Evaluate and bind to the named input")))),Object(c.b)("blockquote",null,Object(c.b)("p",{parentName:"blockquote"},"This is equivalent to providing named arguments to a function in modern programming languages.")),Object(c.b)("h3",{id:"outputs"},"outputs"),Object(c.b)("p",null,"An object for which each key is a variable to assign and the value is one of:"),Object(c.b)("table",null,Object(c.b)("thead",{parentName:"table"},Object(c.b)("tr",{parentName:"thead"},Object(c.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"value"),Object(c.b)("th",Object(n.a)({parentName:"tr"},{align:null}),"meaning"))),Object(c.b)("tbody",{parentName:"table"},Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"null"),Object(c.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Bind variable to output w/ same name (equivalent to ",Object(c.b)("inlineCode",{parentName:"td"},"$(OUTPUT_NAME)"),")")),Object(c.b)("tr",{parentName:"tbody"},Object(c.b)("td",Object(n.a)({parentName:"tr"},{align:null}),Object(c.b)("a",Object(n.a)({parentName:"td"},{href:"/docs/opspec/reference/structure/op-directory/op/reference"}),"reference")),Object(c.b)("td",Object(n.a)({parentName:"tr"},{align:null}),"Bind named output to referenced variable")))),Object(c.b)("blockquote",null,Object(c.b)("p",{parentName:"blockquote"},"This is equivalent to consuming named return values from a function in modern programming languages.")))}u.isMDXComponent=!0},181:function(e,t,r){"use strict";r.d(t,"a",(function(){return u})),r.d(t,"b",(function(){return m}));var n=r(0),a=r.n(n);function c(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function l(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?l(Object(r),!0).forEach((function(t){c(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):l(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function p(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},c=Object.keys(e);for(n=0;n<c.length;n++)r=c[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(e);for(n=0;n<c.length;n++)r=c[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var b=a.a.createContext({}),i=function(e){var t=a.a.useContext(b),r=t;return e&&(r="function"==typeof e?e(t):o({},t,{},e)),r},u=function(e){var t=i(e.components);return a.a.createElement(b.Provider,{value:t},e.children)},d="mdxType",s={inlineCode:"code",wrapper:function(e){var t=e.children;return a.a.createElement(a.a.Fragment,{},t)}},O=Object(n.forwardRef)((function(e,t){var r=e.components,n=e.mdxType,c=e.originalType,l=e.parentName,b=p(e,["components","mdxType","originalType","parentName"]),u=i(r),d=n,O=u["".concat(l,".").concat(d)]||u[d]||s[d]||c;return r?a.a.createElement(O,o({ref:t},b,{components:r})):a.a.createElement(O,o({ref:t},b))}));function m(e,t){var r=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var c=r.length,l=new Array(c);l[0]=O;var o={};for(var p in t)hasOwnProperty.call(t,p)&&(o[p]=t[p]);o.originalType=e,o[d]="string"==typeof e?e:n,l[1]=o;for(var b=2;b<c;b++)l[b]=r[b];return a.a.createElement.apply(null,l)}return a.a.createElement.apply(null,r)}O.displayName="MDXCreateElement"}}]);