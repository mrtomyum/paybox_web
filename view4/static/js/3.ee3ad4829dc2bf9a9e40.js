webpackJsonp([3],{141:function(t,e,n){var r=n(125)(n(184),n(281),null,null,null);t.exports=r.exports},148:function(t,e,n){"use strict";function r(t){return"[object Array]"===C.call(t)}function o(t){return"[object ArrayBuffer]"===C.call(t)}function i(t){return"undefined"!=typeof FormData&&t instanceof FormData}function s(t){return"undefined"!=typeof ArrayBuffer&&ArrayBuffer.isView?ArrayBuffer.isView(t):t&&t.buffer&&t.buffer instanceof ArrayBuffer}function a(t){return"string"==typeof t}function u(t){return"number"==typeof t}function c(t){return void 0===t}function l(t){return null!==t&&"object"==typeof t}function f(t){return"[object Date]"===C.call(t)}function d(t){return"[object File]"===C.call(t)}function m(t){return"[object Blob]"===C.call(t)}function h(t){return"[object Function]"===C.call(t)}function p(t){return l(t)&&h(t.pipe)}function v(t){return"undefined"!=typeof URLSearchParams&&t instanceof URLSearchParams}function y(t){return t.replace(/^\s*/,"").replace(/\s*$/,"")}function g(){return("undefined"==typeof navigator||"ReactNative"!==navigator.product)&&("undefined"!=typeof window&&"undefined"!=typeof document)}function b(t,e){if(null!==t&&void 0!==t)if("object"==typeof t||r(t)||(t=[t]),r(t))for(var n=0,o=t.length;n<o;n++)e.call(null,t[n],n,t);else for(var i in t)Object.prototype.hasOwnProperty.call(t,i)&&e.call(null,t[i],i,t)}function _(){function t(t,n){"object"==typeof e[n]&&"object"==typeof t?e[n]=_(e[n],t):e[n]=t}for(var e={},n=0,r=arguments.length;n<r;n++)b(arguments[n],t);return e}function T(t,e,n){return b(e,function(e,r){t[r]=n&&"function"==typeof e?w(e,n):e}),t}var w=n(154),x=n(177),C=Object.prototype.toString;t.exports={isArray:r,isArrayBuffer:o,isBuffer:x,isFormData:i,isArrayBufferView:s,isString:a,isNumber:u,isObject:l,isUndefined:c,isDate:f,isFile:d,isBlob:m,isFunction:h,isStream:p,isURLSearchParams:v,isStandardBrowserEnv:g,forEach:b,merge:_,extend:T,trim:y}},149:function(t,e,n){"use strict";(function(e){function r(t,e){!o.isUndefined(t)&&o.isUndefined(t["Content-Type"])&&(t["Content-Type"]=e)}var o=n(148),i=n(171),s={"Content-Type":"application/json"},a={adapter:function(){var t;return"undefined"!=typeof XMLHttpRequest?t=n(150):void 0!==e&&(t=n(150)),t}(),transformRequest:[function(t,e){return i(e,"Content-Type"),o.isFormData(t)||o.isArrayBuffer(t)||o.isBuffer(t)||o.isStream(t)||o.isFile(t)||o.isBlob(t)?t:o.isArrayBufferView(t)?t.buffer:o.isURLSearchParams(t)?(r(e,"application/json;charset=utf-8"),t.toString()):o.isObject(t)?(r(e,"application/json;charset=utf-8"),JSON.stringify(t)):t}],transformResponse:[function(t){if("string"==typeof t)try{t=JSON.parse(t)}catch(t){}return t}],timeout:0,xsrfCookieName:"XSRF-TOKEN",xsrfHeaderName:"X-XSRF-TOKEN",maxContentLength:-1,validateStatus:function(t){return t>=200&&t<300}};a.headers={common:{Accept:"application/json, */*","Content-Type":"application/json"}},o.forEach(["delete","get","head"],function(t){a.headers[t]={}}),o.forEach(["post","put","patch"],function(t){a.headers[t]=o.merge(s)}),t.exports=a}).call(e,n(178))},150:function(t,e,n){"use strict";var r=n(148),o=n(163),i=n(166),s=n(172),a=n(170),u=n(153),c="undefined"!=typeof window&&window.btoa&&window.btoa.bind(window)||n(165);t.exports=function(t){return new Promise(function(e,l){var f=t.data,d=t.headers;r.isFormData(f)&&delete d["Content-Type"];var m=new XMLHttpRequest,h="onreadystatechange",p=!1;if("undefined"==typeof window||!window.XDomainRequest||"withCredentials"in m||a(t.url)||(m=new window.XDomainRequest,h="onload",p=!0,m.onprogress=function(){},m.ontimeout=function(){}),t.auth){var v=t.auth.username||"",y=t.auth.password||"";d.Authorization="Basic "+c(v+":"+y)}if(m.open(t.method.toUpperCase(),i(t.url,t.params,t.paramsSerializer),!0),m.timeout=t.timeout,m[h]=function(){if(m&&(4===m.readyState||p)&&(0!==m.status||m.responseURL&&0===m.responseURL.indexOf("file:"))){var n="getAllResponseHeaders"in m?s(m.getAllResponseHeaders()):null,r=t.responseType&&"text"!==t.responseType?m.response:m.responseText,i={data:r,status:1223===m.status?204:m.status,statusText:1223===m.status?"No Content":m.statusText,headers:n,config:t,request:m};o(e,l,i),m=null}},m.onerror=function(){l(u("Network Error",t,null,m)),m=null},m.ontimeout=function(){l(u("timeout of "+t.timeout+"ms exceeded",t,"ECONNABORTED",m)),m=null},r.isStandardBrowserEnv()){var g=n(168),b=(t.withCredentials||a(t.url))&&t.xsrfCookieName?g.read(t.xsrfCookieName):void 0;b&&(d[t.xsrfHeaderName]=b)}if("setRequestHeader"in m&&r.forEach(d,function(t,e){void 0===f&&"content-type"===e.toLowerCase()?delete d[e]:m.setRequestHeader(e,t)}),t.withCredentials&&(m.withCredentials=!0),t.responseType)try{m.responseType=t.responseType}catch(e){if("json"!==t.responseType)throw e}"function"==typeof t.onDownloadProgress&&m.addEventListener("progress",t.onDownloadProgress),"function"==typeof t.onUploadProgress&&m.upload&&m.upload.addEventListener("progress",t.onUploadProgress),t.cancelToken&&t.cancelToken.promise.then(function(t){m&&(m.abort(),l(t),m=null)}),void 0===f&&(f=null),m.send(f)})}},151:function(t,e,n){"use strict";function r(t){this.message=t}r.prototype.toString=function(){return"Cancel"+(this.message?": "+this.message:"")},r.prototype.__CANCEL__=!0,t.exports=r},152:function(t,e,n){"use strict";t.exports=function(t){return!(!t||!t.__CANCEL__)}},153:function(t,e,n){"use strict";var r=n(162);t.exports=function(t,e,n,o,i){var s=new Error(t);return r(s,e,n,o,i)}},154:function(t,e,n){"use strict";t.exports=function(t,e){return function(){for(var n=new Array(arguments.length),r=0;r<n.length;r++)n[r]=arguments[r];return t.apply(e,n)}}},155:function(t,e,n){t.exports={default:n(175),__esModule:!0}},156:function(t,e,n){t.exports=n(157)},157:function(t,e,n){"use strict";function r(t){var e=new s(t),n=i(s.prototype.request,e);return o.extend(n,s.prototype,e),o.extend(n,e),n}var o=n(148),i=n(154),s=n(159),a=n(149),u=r(a);u.Axios=s,u.create=function(t){return r(o.merge(a,t))},u.Cancel=n(151),u.CancelToken=n(158),u.isCancel=n(152),u.all=function(t){return Promise.all(t)},u.spread=n(173),t.exports=u,t.exports.default=u},158:function(t,e,n){"use strict";function r(t){if("function"!=typeof t)throw new TypeError("executor must be a function.");var e;this.promise=new Promise(function(t){e=t});var n=this;t(function(t){n.reason||(n.reason=new o(t),e(n.reason))})}var o=n(151);r.prototype.throwIfRequested=function(){if(this.reason)throw this.reason},r.source=function(){var t;return{token:new r(function(e){t=e}),cancel:t}},t.exports=r},159:function(t,e,n){"use strict";function r(t){this.defaults=t,this.interceptors={request:new s,response:new s}}var o=n(149),i=n(148),s=n(160),a=n(161),u=n(169),c=n(167);r.prototype.request=function(t){"string"==typeof t&&(t=i.merge({url:arguments[0]},arguments[1])),t=i.merge(o,this.defaults,{method:"get"},t),t.method=t.method.toLowerCase(),t.baseURL&&!u(t.url)&&(t.url=c(t.baseURL,t.url));var e=[a,void 0],n=Promise.resolve(t);for(this.interceptors.request.forEach(function(t){e.unshift(t.fulfilled,t.rejected)}),this.interceptors.response.forEach(function(t){e.push(t.fulfilled,t.rejected)});e.length;)n=n.then(e.shift(),e.shift());return n},i.forEach(["delete","get","head","options"],function(t){r.prototype[t]=function(e,n){return this.request(i.merge(n||{},{method:t,url:e}))}}),i.forEach(["post","put","patch"],function(t){r.prototype[t]=function(e,n,r){return this.request(i.merge(r||{},{method:t,url:e,data:n}))}}),t.exports=r},160:function(t,e,n){"use strict";function r(){this.handlers=[]}var o=n(148);r.prototype.use=function(t,e){return this.handlers.push({fulfilled:t,rejected:e}),this.handlers.length-1},r.prototype.eject=function(t){this.handlers[t]&&(this.handlers[t]=null)},r.prototype.forEach=function(t){o.forEach(this.handlers,function(e){null!==e&&t(e)})},t.exports=r},161:function(t,e,n){"use strict";function r(t){t.cancelToken&&t.cancelToken.throwIfRequested()}var o=n(148),i=n(164),s=n(152),a=n(149);t.exports=function(t){return r(t),t.headers=t.headers||{},t.data=i(t.data,t.headers,t.transformRequest),t.headers=o.merge(t.headers.common||{},t.headers[t.method]||{},t.headers||{}),o.forEach(["delete","get","head","post","put","patch","common"],function(e){delete t.headers[e]}),(t.adapter||a.adapter)(t).then(function(e){return r(t),e.data=i(e.data,e.headers,t.transformResponse),e},function(e){return s(e)||(r(t),e&&e.response&&(e.response.data=i(e.response.data,e.response.headers,t.transformResponse))),Promise.reject(e)})}},162:function(t,e,n){"use strict";t.exports=function(t,e,n,r,o){return t.config=e,n&&(t.code=n),t.request=r,t.response=o,t}},163:function(t,e,n){"use strict";var r=n(153);t.exports=function(t,e,n){var o=n.config.validateStatus;n.status&&o&&!o(n.status)?e(r("Request failed with status code "+n.status,n.config,null,n.request,n)):t(n)}},164:function(t,e,n){"use strict";var r=n(148);t.exports=function(t,e,n){return r.forEach(n,function(n){t=n(t,e)}),t}},165:function(t,e,n){"use strict";function r(){this.message="String contains an invalid character"}function o(t){for(var e,n,o=String(t),s="",a=0,u=i;o.charAt(0|a)||(u="=",a%1);s+=u.charAt(63&e>>8-a%1*8)){if((n=o.charCodeAt(a+=.75))>255)throw new r;e=e<<8|n}return s}var i="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";r.prototype=new Error,r.prototype.code=5,r.prototype.name="InvalidCharacterError",t.exports=o},166:function(t,e,n){"use strict";function r(t){return encodeURIComponent(t).replace(/%40/gi,"@").replace(/%3A/gi,":").replace(/%24/g,"$").replace(/%2C/gi,",").replace(/%20/g,"+").replace(/%5B/gi,"[").replace(/%5D/gi,"]")}var o=n(148);t.exports=function(t,e,n){if(!e)return t;var i;if(n)i=n(e);else if(o.isURLSearchParams(e))i=e.toString();else{var s=[];o.forEach(e,function(t,e){null!==t&&void 0!==t&&(o.isArray(t)&&(e+="[]"),o.isArray(t)||(t=[t]),o.forEach(t,function(t){o.isDate(t)?t=t.toISOString():o.isObject(t)&&(t=JSON.stringify(t)),s.push(r(e)+"="+r(t))}))}),i=s.join("&")}return i&&(t+=(-1===t.indexOf("?")?"?":"&")+i),t}},167:function(t,e,n){"use strict";t.exports=function(t,e){return e?t.replace(/\/+$/,"")+"/"+e.replace(/^\/+/,""):t}},168:function(t,e,n){"use strict";var r=n(148);t.exports=r.isStandardBrowserEnv()?function(){return{write:function(t,e,n,o,i,s){var a=[];a.push(t+"="+encodeURIComponent(e)),r.isNumber(n)&&a.push("expires="+new Date(n).toGMTString()),r.isString(o)&&a.push("path="+o),r.isString(i)&&a.push("domain="+i),!0===s&&a.push("secure"),document.cookie=a.join("; ")},read:function(t){var e=document.cookie.match(new RegExp("(^|;\\s*)("+t+")=([^;]*)"));return e?decodeURIComponent(e[3]):null},remove:function(t){this.write(t,"",Date.now()-864e5)}}}():function(){return{write:function(){},read:function(){return null},remove:function(){}}}()},169:function(t,e,n){"use strict";t.exports=function(t){return/^([a-z][a-z\d\+\-\.]*:)?\/\//i.test(t)}},170:function(t,e,n){"use strict";var r=n(148);t.exports=r.isStandardBrowserEnv()?function(){function t(t){var e=t;return n&&(o.setAttribute("href",e),e=o.href),o.setAttribute("href",e),{href:o.href,protocol:o.protocol?o.protocol.replace(/:$/,""):"",host:o.host,search:o.search?o.search.replace(/^\?/,""):"",hash:o.hash?o.hash.replace(/^#/,""):"",hostname:o.hostname,port:o.port,pathname:"/"===o.pathname.charAt(0)?o.pathname:"/"+o.pathname}}var e,n=/(msie|trident)/i.test(navigator.userAgent),o=document.createElement("a");return e=t(window.location.href),function(n){var o=r.isString(n)?t(n):n;return o.protocol===e.protocol&&o.host===e.host}}():function(){return function(){return!0}}()},171:function(t,e,n){"use strict";var r=n(148);t.exports=function(t,e){r.forEach(t,function(n,r){r!==e&&r.toUpperCase()===e.toUpperCase()&&(t[e]=n,delete t[r])})}},172:function(t,e,n){"use strict";var r=n(148);t.exports=function(t){var e,n,o,i={};return t?(r.forEach(t.split("\n"),function(t){o=t.indexOf(":"),e=r.trim(t.substr(0,o)).toLowerCase(),n=r.trim(t.substr(o+1)),e&&(i[e]=i[e]?i[e]+", "+n:n)}),i):i}},173:function(t,e,n){"use strict";t.exports=function(t){return function(e){return t.apply(null,e)}}},174:function(t,e,n){"use strict";var r=n(155),o=n.n(r),i=n(1),s=n(156),a=n.n(s),u=n(179),c=n.n(u);i.a.use(c.a,a.a),a.a.defaults.headers.post["Content-Type"]="application/json";var l="http://localhost";e.a={getMenuAX:function(t,e){i.a.axios.get(l+":8888/menu").then(function(e){t(e.data)},function(t){e(t)})},getItemAX:function(t,e,n){i.a.axios.get(l+":8888/menu/"+t+"/").then(function(t){e(t.data)},function(t){n(t)})},sendOrderAX:function(t,e,n){console.log(o()(t));i.a.axios.post("http://localhost:8888/sale",o()(t)).then(function(t){e(t.data)},function(t){n(t)})},getMoneyOnMachine:function(t,e){i.a.axios.get(l+":8888/cash").then(function(e){t(e.data)},function(t){e(t)})}}},175:function(t,e,n){var r=n(176),o=r.JSON||(r.JSON={stringify:JSON.stringify});t.exports=function(t){return o.stringify.apply(o,arguments)}},176:function(t,e){var n=t.exports={version:"2.5.0"};"number"==typeof __e&&(__e=n)},177:function(t,e){function n(t){return!!t.constructor&&"function"==typeof t.constructor.isBuffer&&t.constructor.isBuffer(t)}function r(t){return"function"==typeof t.readFloatLE&&"function"==typeof t.slice&&n(t.slice(0,0))}/*!
 * Determine if an object is a Buffer
 *
 * @author   Feross Aboukhadijeh <feross@feross.org> <http://feross.org>
 * @license  MIT
 */
t.exports=function(t){return null!=t&&(n(t)||r(t)||!!t._isBuffer)}},178:function(t,e){function n(){throw new Error("setTimeout has not been defined")}function r(){throw new Error("clearTimeout has not been defined")}function o(t){if(l===setTimeout)return setTimeout(t,0);if((l===n||!l)&&setTimeout)return l=setTimeout,setTimeout(t,0);try{return l(t,0)}catch(e){try{return l.call(null,t,0)}catch(e){return l.call(this,t,0)}}}function i(t){if(f===clearTimeout)return clearTimeout(t);if((f===r||!f)&&clearTimeout)return f=clearTimeout,clearTimeout(t);try{return f(t)}catch(e){try{return f.call(null,t)}catch(e){return f.call(this,t)}}}function s(){p&&m&&(p=!1,m.length?h=m.concat(h):v=-1,h.length&&a())}function a(){if(!p){var t=o(s);p=!0;for(var e=h.length;e;){for(m=h,h=[];++v<e;)m&&m[v].run();v=-1,e=h.length}m=null,p=!1,i(t)}}function u(t,e){this.fun=t,this.array=e}function c(){}var l,f,d=t.exports={};!function(){try{l="function"==typeof setTimeout?setTimeout:n}catch(t){l=n}try{f="function"==typeof clearTimeout?clearTimeout:r}catch(t){f=r}}();var m,h=[],p=!1,v=-1;d.nextTick=function(t){var e=new Array(arguments.length-1);if(arguments.length>1)for(var n=1;n<arguments.length;n++)e[n-1]=arguments[n];h.push(new u(t,e)),1!==h.length||p||o(a)},u.prototype.run=function(){this.fun.apply(null,this.array)},d.title="browser",d.browser=!0,d.env={},d.argv=[],d.version="",d.versions={},d.on=c,d.addListener=c,d.once=c,d.off=c,d.removeListener=c,d.removeAllListeners=c,d.emit=c,d.prependListener=c,d.prependOnceListener=c,d.listeners=function(t){return[]},d.binding=function(t){throw new Error("process.binding is not supported")},d.cwd=function(){return"/"},d.chdir=function(t){throw new Error("process.chdir is not supported")},d.umask=function(){return 0}},179:function(t,e,n){"use strict";var r,o,i="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(t){return typeof t}:function(t){return t&&"function"==typeof Symbol&&t.constructor===Symbol&&t!==Symbol.prototype?"symbol":typeof t};!function(){function n(t,e){if(!n.installed){if(n.installed=!0,!e)return void console.error("You have to install axios");t.axios=e,Object.defineProperties(t.prototype,{axios:{get:function(){return e}},$http:{get:function(){return e}}})}}"object"==i(e)?t.exports=n:(r=[],void 0!==(o=function(){return n}.apply(e,r))&&(t.exports=o))}()},180:function(t,e,n){var r,o;/*! @preserve
 * numeral.js
 * version : 2.0.6
 * author : Adam Draper
 * license : MIT
 * http://adamwdraper.github.com/Numeral-js/
 */
!function(i,s){r=s,void 0!==(o="function"==typeof r?r.call(e,n,e,t):r)&&(t.exports=o)}(0,function(){function t(t,e){this._input=t,this._value=e}var e,n,r={},o={},i={currentLocale:"en",zeroFormat:null,nullFormat:null,defaultFormat:"0,0",scalePercentBy100:!0},s={currentLocale:i.currentLocale,zeroFormat:i.zeroFormat,nullFormat:i.nullFormat,defaultFormat:i.defaultFormat,scalePercentBy100:i.scalePercentBy100};return e=function(o){var i,a,u,c;if(e.isNumeral(o))i=o.value();else if(0===o||void 0===o)i=0;else if(null===o||n.isNaN(o))i=null;else if("string"==typeof o)if(s.zeroFormat&&o===s.zeroFormat)i=0;else if(s.nullFormat&&o===s.nullFormat||!o.replace(/[^0-9]+/g,"").length)i=null;else{for(a in r)if((c="function"==typeof r[a].regexps.unformat?r[a].regexps.unformat():r[a].regexps.unformat)&&o.match(c)){u=r[a].unformat;break}u=u||e._.stringToNumber,i=u(o)}else i=Number(o)||null;return new t(o,i)},e.version="2.0.6",e.isNumeral=function(e){return e instanceof t},e._=n={numberToFormat:function(t,n,r){var i,s,a,u,c,l,f,d=o[e.options.currentLocale],m=!1,h=!1,p=0,v="",y="",g=!1;if(t=t||0,s=Math.abs(t),e._.includes(n,"(")?(m=!0,n=n.replace(/[\(|\)]/g,"")):(e._.includes(n,"+")||e._.includes(n,"-"))&&(c=e._.includes(n,"+")?n.indexOf("+"):t<0?n.indexOf("-"):-1,n=n.replace(/[\+|\-]/g,"")),e._.includes(n,"a")&&(i=n.match(/a(k|m|b|t)?/),i=!!i&&i[1],e._.includes(n," a")&&(v=" "),n=n.replace(new RegExp(v+"a[kmbt]?"),""),s>=1e12&&!i||"t"===i?(v+=d.abbreviations.trillion,t/=1e12):s<1e12&&s>=1e9&&!i||"b"===i?(v+=d.abbreviations.billion,t/=1e9):s<1e9&&s>=1e6&&!i||"m"===i?(v+=d.abbreviations.million,t/=1e6):(s<1e6&&s>=1e3&&!i||"k"===i)&&(v+=d.abbreviations.thousand,t/=1e3)),e._.includes(n,"[.]")&&(h=!0,n=n.replace("[.]",".")),a=t.toString().split(".")[0],u=n.split(".")[1],l=n.indexOf(","),p=(n.split(".")[0].split(",")[0].match(/0/g)||[]).length,u?(e._.includes(u,"[")?(u=u.replace("]",""),u=u.split("["),y=e._.toFixed(t,u[0].length+u[1].length,r,u[1].length)):y=e._.toFixed(t,u.length,r),a=y.split(".")[0],y=e._.includes(y,".")?d.delimiters.decimal+y.split(".")[1]:"",h&&0===Number(y.slice(1))&&(y="")):a=e._.toFixed(t,0,r),v&&!i&&Number(a)>=1e3&&v!==d.abbreviations.trillion)switch(a=String(Number(a)/1e3),v){case d.abbreviations.thousand:v=d.abbreviations.million;break;case d.abbreviations.million:v=d.abbreviations.billion;break;case d.abbreviations.billion:v=d.abbreviations.trillion}if(e._.includes(a,"-")&&(a=a.slice(1),g=!0),a.length<p)for(var b=p-a.length;b>0;b--)a="0"+a;return l>-1&&(a=a.toString().replace(/(\d)(?=(\d{3})+(?!\d))/g,"$1"+d.delimiters.thousands)),0===n.indexOf(".")&&(a=""),f=a+y+(v||""),m?f=(m&&g?"(":"")+f+(m&&g?")":""):c>=0?f=0===c?(g?"-":"+")+f:f+(g?"-":"+"):g&&(f="-"+f),f},stringToNumber:function(t){var e,n,r,i=o[s.currentLocale],a=t,u={thousand:3,million:6,billion:9,trillion:12};if(s.zeroFormat&&t===s.zeroFormat)n=0;else if(s.nullFormat&&t===s.nullFormat||!t.replace(/[^0-9]+/g,"").length)n=null;else{n=1,"."!==i.delimiters.decimal&&(t=t.replace(/\./g,"").replace(i.delimiters.decimal,"."));for(e in u)if(r=new RegExp("[^a-zA-Z]"+i.abbreviations[e]+"(?:\\)|(\\"+i.currency.symbol+")?(?:\\))?)?$"),a.match(r)){n*=Math.pow(10,u[e]);break}n*=(t.split("-").length+Math.min(t.split("(").length-1,t.split(")").length-1))%2?1:-1,t=t.replace(/[^0-9\.]+/g,""),n*=Number(t)}return n},isNaN:function(t){return"number"==typeof t&&isNaN(t)},includes:function(t,e){return-1!==t.indexOf(e)},insert:function(t,e,n){return t.slice(0,n)+e+t.slice(n)},reduce:function(t,e){if(null===this)throw new TypeError("Array.prototype.reduce called on null or undefined");if("function"!=typeof e)throw new TypeError(e+" is not a function");var n,r=Object(t),o=r.length>>>0,i=0;if(3===arguments.length)n=arguments[2];else{for(;i<o&&!(i in r);)i++;if(i>=o)throw new TypeError("Reduce of empty array with no initial value");n=r[i++]}for(;i<o;i++)i in r&&(n=e(n,r[i],i,r));return n},multiplier:function(t){var e=t.toString().split(".");return e.length<2?1:Math.pow(10,e[1].length)},correctionFactor:function(){return Array.prototype.slice.call(arguments).reduce(function(t,e){var r=n.multiplier(e);return t>r?t:r},1)},toFixed:function(t,e,n,r){var o,i,s,a,u=t.toString().split("."),c=e-(r||0);return o=2===u.length?Math.min(Math.max(u[1].length,c),e):c,s=Math.pow(10,o),a=(n(t+"e+"+o)/s).toFixed(o),r>e-o&&(i=new RegExp("\\.?0{1,"+(r-(e-o))+"}$"),a=a.replace(i,"")),a}},e.options=s,e.formats=r,e.locales=o,e.locale=function(t){return t&&(s.currentLocale=t.toLowerCase()),s.currentLocale},e.localeData=function(t){if(!t)return o[s.currentLocale];if(t=t.toLowerCase(),!o[t])throw new Error("Unknown locale : "+t);return o[t]},e.reset=function(){for(var t in i)s[t]=i[t]},e.zeroFormat=function(t){s.zeroFormat="string"==typeof t?t:null},e.nullFormat=function(t){s.nullFormat="string"==typeof t?t:null},e.defaultFormat=function(t){s.defaultFormat="string"==typeof t?t:"0.0"},e.register=function(t,e,n){if(e=e.toLowerCase(),this[t+"s"][e])throw new TypeError(e+" "+t+" already registered.");return this[t+"s"][e]=n,n},e.validate=function(t,n){var r,o,i,s,a,u,c,l;if("string"!=typeof t&&(t+="",console.warn&&console.warn("Numeral.js: Value is not string. It has been co-erced to: ",t)),t=t.trim(),t.match(/^\d+$/))return!0;if(""===t)return!1;try{c=e.localeData(n)}catch(t){c=e.localeData(e.locale())}return i=c.currency.symbol,a=c.abbreviations,r=c.delimiters.decimal,o="."===c.delimiters.thousands?"\\.":c.delimiters.thousands,(null===(l=t.match(/^[^\d]+/))||(t=t.substr(1),l[0]===i))&&((null===(l=t.match(/[^\d]+$/))||(t=t.slice(0,-1),l[0]===a.thousand||l[0]===a.million||l[0]===a.billion||l[0]===a.trillion))&&(u=new RegExp(o+"{2}"),!t.match(/[^\d.,]/g)&&(s=t.split(r),!(s.length>2)&&(s.length<2?!!s[0].match(/^\d+.*\d$/)&&!s[0].match(u):1===s[0].length?!!s[0].match(/^\d+$/)&&!s[0].match(u)&&!!s[1].match(/^\d+$/):!!s[0].match(/^\d+.*\d$/)&&!s[0].match(u)&&!!s[1].match(/^\d+$/)))))},e.fn=t.prototype={clone:function(){return e(this)},format:function(t,n){var o,i,a,u=this._value,c=t||s.defaultFormat;if(n=n||Math.round,0===u&&null!==s.zeroFormat)i=s.zeroFormat;else if(null===u&&null!==s.nullFormat)i=s.nullFormat;else{for(o in r)if(c.match(r[o].regexps.format)){a=r[o].format;break}a=a||e._.numberToFormat,i=a(u,c,n)}return i},value:function(){return this._value},input:function(){return this._input},set:function(t){return this._value=Number(t),this},add:function(t){function e(t,e,n,o){return t+Math.round(r*e)}var r=n.correctionFactor.call(null,this._value,t);return this._value=n.reduce([this._value,t],e,0)/r,this},subtract:function(t){function e(t,e,n,o){return t-Math.round(r*e)}var r=n.correctionFactor.call(null,this._value,t);return this._value=n.reduce([t],e,Math.round(this._value*r))/r,this},multiply:function(t){function e(t,e,r,o){var i=n.correctionFactor(t,e);return Math.round(t*i)*Math.round(e*i)/Math.round(i*i)}return this._value=n.reduce([this._value,t],e,1),this},divide:function(t){function e(t,e,r,o){var i=n.correctionFactor(t,e);return Math.round(t*i)/Math.round(e*i)}return this._value=n.reduce([this._value,t],e),this},difference:function(t){return Math.abs(e(this._value).subtract(t).value())}},e.register("locale","en",{delimiters:{thousands:",",decimal:"."},abbreviations:{thousand:"k",million:"m",billion:"b",trillion:"t"},ordinal:function(t){var e=t%10;return 1==~~(t%100/10)?"th":1===e?"st":2===e?"nd":3===e?"rd":"th"},currency:{symbol:"$"}}),function(){e.register("format","bps",{regexps:{format:/(BPS)/,unformat:/(BPS)/},format:function(t,n,r){var o,i=e._.includes(n," BPS")?" ":"";return t*=1e4,n=n.replace(/\s?BPS/,""),o=e._.numberToFormat(t,n,r),e._.includes(o,")")?(o=o.split(""),o.splice(-1,0,i+"BPS"),o=o.join("")):o=o+i+"BPS",o},unformat:function(t){return+(1e-4*e._.stringToNumber(t)).toFixed(15)}})}(),function(){var t={base:1e3,suffixes:["B","KB","MB","GB","TB","PB","EB","ZB","YB"]},n={base:1024,suffixes:["B","KiB","MiB","GiB","TiB","PiB","EiB","ZiB","YiB"]},r=t.suffixes.concat(n.suffixes.filter(function(e){return t.suffixes.indexOf(e)<0})),o=r.join("|");o="("+o.replace("B","B(?!PS)")+")",e.register("format","bytes",{regexps:{format:/([0\s]i?b)/,unformat:new RegExp(o)},format:function(r,o,i){var s,a,u,c=e._.includes(o,"ib")?n:t,l=e._.includes(o," b")||e._.includes(o," ib")?" ":"";for(o=o.replace(/\s?i?b/,""),s=0;s<=c.suffixes.length;s++)if(a=Math.pow(c.base,s),u=Math.pow(c.base,s+1),null===r||0===r||r>=a&&r<u){l+=c.suffixes[s],a>0&&(r/=a);break}return e._.numberToFormat(r,o,i)+l},unformat:function(r){var o,i,s=e._.stringToNumber(r);if(s){for(o=t.suffixes.length-1;o>=0;o--){if(e._.includes(r,t.suffixes[o])){i=Math.pow(t.base,o);break}if(e._.includes(r,n.suffixes[o])){i=Math.pow(n.base,o);break}}s*=i||1}return s}})}(),function(){e.register("format","currency",{regexps:{format:/(\$)/},format:function(t,n,r){var o,i,s=e.locales[e.options.currentLocale],a={before:n.match(/^([\+|\-|\(|\s|\$]*)/)[0],after:n.match(/([\+|\-|\)|\s|\$]*)$/)[0]};for(n=n.replace(/\s?\$\s?/,""),o=e._.numberToFormat(t,n,r),t>=0?(a.before=a.before.replace(/[\-\(]/,""),a.after=a.after.replace(/[\-\)]/,"")):t<0&&!e._.includes(a.before,"-")&&!e._.includes(a.before,"(")&&(a.before="-"+a.before),i=0;i<a.before.length;i++)switch(a.before[i]){case"$":o=e._.insert(o,s.currency.symbol,i);break;case" ":o=e._.insert(o," ",i+s.currency.symbol.length-1)}for(i=a.after.length-1;i>=0;i--)switch(a.after[i]){case"$":o=i===a.after.length-1?o+s.currency.symbol:e._.insert(o,s.currency.symbol,-(a.after.length-(1+i)));break;case" ":o=i===a.after.length-1?o+" ":e._.insert(o," ",-(a.after.length-(1+i)+s.currency.symbol.length-1))}return o}})}(),function(){e.register("format","exponential",{regexps:{format:/(e\+|e-)/,unformat:/(e\+|e-)/},format:function(t,n,r){var o="number"!=typeof t||e._.isNaN(t)?"0e+0":t.toExponential(),i=o.split("e");return n=n.replace(/e[\+|\-]{1}0/,""),e._.numberToFormat(Number(i[0]),n,r)+"e"+i[1]},unformat:function(t){function n(t,n,r,o){var i=e._.correctionFactor(t,n);return t*i*(n*i)/(i*i)}var r=e._.includes(t,"e+")?t.split("e+"):t.split("e-"),o=Number(r[0]),i=Number(r[1]);return i=e._.includes(t,"e-")?i*=-1:i,e._.reduce([o,Math.pow(10,i)],n,1)}})}(),function(){e.register("format","ordinal",{regexps:{format:/(o)/},format:function(t,n,r){var o=e.locales[e.options.currentLocale],i=e._.includes(n," o")?" ":"";return n=n.replace(/\s?o/,""),i+=o.ordinal(t),e._.numberToFormat(t,n,r)+i}})}(),function(){e.register("format","percentage",{regexps:{format:/(%)/,unformat:/(%)/},format:function(t,n,r){var o,i=e._.includes(n," %")?" ":"";return e.options.scalePercentBy100&&(t*=100),n=n.replace(/\s?\%/,""),o=e._.numberToFormat(t,n,r),e._.includes(o,")")?(o=o.split(""),o.splice(-1,0,i+"%"),o=o.join("")):o=o+i+"%",o},unformat:function(t){var n=e._.stringToNumber(t);return e.options.scalePercentBy100?.01*n:n}})}(),function(){e.register("format","time",{regexps:{format:/(:)/,unformat:/(:)/},format:function(t,e,n){var r=Math.floor(t/60/60),o=Math.floor((t-60*r*60)/60),i=Math.round(t-60*r*60-60*o);return r+":"+(o<10?"0"+o:o)+":"+(i<10?"0"+i:i)},unformat:function(t){var e=t.split(":"),n=0;return 3===e.length?(n+=60*Number(e[0])*60,n+=60*Number(e[1]),n+=Number(e[2])):2===e.length&&(n+=60*Number(e[0]),n+=Number(e[1])),Number(n)}})}(),e})},181:function(t,e,n){t.exports=n.p+"static/img/c1.afead32.png"},182:function(t,e,n){t.exports=n.p+"static/img/c10.0cf6c3a.png"},183:function(t,e,n){t.exports=n.p+"static/img/c5.3693c5b.png"},184:function(t,e,n){"use strict";Object.defineProperty(e,"__esModule",{value:!0});var r=n(155),o=n.n(r),i=n(180),s=n.n(i);n(174);e.default={name:"add_money",data:function(){return{titlePage:"เพิ่มยอดเงินในตู้",QT1:0,QT2:0,QT3:0,QT4:0,onMachine:0,addMachine:0,v1:0,v2:0,v5:0,v10:0,v100:0}},methods:{add_QTY:function(t){switch(this.stopSound(),this.Soundclick(),t){case 0:this.QT1<10?(this.QT1+=1,this.addMachine+=1e3,this.v1+=1e3):alert("ท่านใส่ยอดเงินมากเกินไป!"),document.getElementsByClassName("add-QTY")[t].style.display="block",document.getElementsByClassName("rm-QTY")[t].style.display="block";break;case 1:this.QT2<10?(this.QT2+=1,this.addMachine+=1e3,this.v5+=1e3):alert("ท่านใส่ยอดเงินมากเกินไป!"),document.getElementsByClassName("add-QTY")[t].style.display="block",document.getElementsByClassName("rm-QTY")[t].style.display="block";break;case 2:this.QT3<10?(this.QT3+=1,this.addMachine+=1e3,this.v10+=1e3):alert("ท่านใส่ยอดเงินมากเกินไป!"),document.getElementsByClassName("add-QTY")[t].style.display="block",document.getElementsByClassName("rm-QTY")[t].style.display="block";break;case 3:this.QT4<10?(this.QT4+=1,this.addMachine+=1e3,this.v100+=1e3):alert("ท่านใส่ยอดเงินมากเกินไป!"),document.getElementsByClassName("add-QTY")[t].style.display="block",document.getElementsByClassName("rm-QTY")[t].style.display="block"}},rm_QTY:function(t){switch(this.stopSound(),this.Soundclick(),t){case 0:0!=this.QT1&&(this.QT1-=1,this.addMachine-=1e3,this.v1-=1e3),0==this.QT1&&(document.getElementsByClassName("add-QTY")[t].style.display="none",document.getElementsByClassName("rm-QTY")[t].style.display="none");break;case 1:0!=this.QT2&&(this.QT2-=1,this.addMachine-=1e3,this.v5-=1e3),0==this.QT2&&(document.getElementsByClassName("add-QTY")[t].style.display="none",document.getElementsByClassName("rm-QTY")[t].style.display="none");break;case 2:0!=this.QT3&&(this.QT3-=1,this.addMachine-=1e3,this.v10-=1e3),0==this.QT3&&(document.getElementsByClassName("add-QTY")[t].style.display="none",document.getElementsByClassName("rm-QTY")[t].style.display="none");break;case 3:0!=this.QT4&&(this.QT4-=1,this.addMachine-=1e3,this.v100-=1e3),0==this.QT4&&(document.getElementsByClassName("add-QTY")[t].style.display="none",document.getElementsByClassName("rm-QTY")[t].style.display="none")}},backSetting:function(){this.$router.push("/setting"),this.QT1=0,this.QT2=0,this.QT3=0,this.QT4=0,this.addMachine=0;for(var t=0;t<4;t++)document.getElementsByClassName("add-QTY")[t].style.display="none",document.getElementsByClassName("rm-QTY")[t].style.display="none"},format_money:function(t){return s()(t).format("0,0.00")},save_data:function(){this.stopSound(),this.Soundclick(),this.QT1=0,this.QT2=0,this.QT3=0,this.QT4=0,this.addMachine=0;for(var t=0;t<4;t++)document.getElementsByClassName("add-QTY")[t].style.display="none",document.getElementsByClassName("rm-QTY")[t].style.display="none";var e={v025:0,v050:0,v1:this.v1,v2:0,v5:this.v5,v10:this.v10,v100:this.v100};this.$socket.sendObj({Device:"host",type:"request",command:"cash_refill",data:e})},websocket_onmessage:function(){var t=this;this.$options.sockets.onmessage=function(e){return t.event_sk(JSON.parse(e.data))}},event_sk:function(t){switch(console.log(o()(t)),t.command){case"cash_count":case"cash_refill":this.getMoneyOnMachine(t.data)}},getMoneyOnMachine:function(t){this.onMachine=t},Soundclick:function(){document.getElementById("audio").play()},stopSound:function(){document.getElementById("audio").currentTime=0}},beforeDestroy:function(){return{sockets:null}},mounted:function(){this.websocket_onmessage(),setTimeout(function(){this.$socket.sendObj({Device:"host",type:"request",command:"cash_count"})}.bind(this),200)}}},271:function(t,e,n){t.exports=n.p+"static/img/b100.bc057e4.jpg"},281:function(t,e,n){t.exports={render:function(){var t=this,e=t.$createElement,r=t._self._c||e;return r("div",{staticClass:"add_money"},[r("button",{staticClass:"back",on:{click:t.backSetting}},[t._m(0)]),t._v(" "),r("div",{staticClass:"container",staticStyle:{padding:"0 1%"}},[r("h1",{staticStyle:{"font-size":"50px","margin-top":"1%"}},[t._v(" "+t._s(t.titlePage))]),t._v(" "),r("div",{staticClass:"add-l"},[r("h1",[t._v("ยอดเงินในตู้")]),t._v(" "),r("div",{staticClass:"text-money"},[t._v("\n\t\t\t\t\t"+t._s(t.format_money(t.onMachine))+"\n\t\t\t\t")]),t._v(" "),r("h1",[t._v("ยอดเงินที่เพิ่ม")]),t._v(" "),r("div",{staticClass:"text-money"},[t._v("\n\t\t\t\t\t"+t._s(t.format_money(t.addMachine))+"\n\t\t\t\t")]),t._v(" "),r("button",{staticClass:"button is-success button-money",on:{click:t.save_data}},[t._v("\n\t\t\t\t\tบันทึก\n\t\t\t\t")])]),t._v(" "),r("div",{staticClass:"add-r"},[r("div",{staticClass:"data-money"},[r("div",{staticClass:"add-QTY"},[t._v(t._s(t.QT1))]),t._v(" "),r("button",{staticClass:"bt_money",on:{click:function(e){t.add_QTY(0)}}},[r("img",{attrs:{src:n(181)}}),t._v(" "),r("h1",[t._v("เหรียญ 1 บาท")])]),t._v(" "),r("div",{staticClass:"rm-QTY",on:{click:function(e){t.rm_QTY(0)}}},[t._v(" - ")])]),t._v(" "),r("div",{staticClass:"data-money"},[r("div",{staticClass:"add-QTY"},[t._v(t._s(t.QT2))]),t._v(" "),r("button",{staticClass:"bt_money",on:{click:function(e){t.add_QTY(1)}}},[r("img",{attrs:{src:n(183)}}),t._v(" "),r("h1",[t._v("เหรียญ 5 บาท")])]),t._v(" "),r("div",{staticClass:"rm-QTY",on:{click:function(e){t.rm_QTY(1)}}},[t._v(" - ")])]),t._v(" "),r("div",{staticClass:"data-money"},[r("div",{staticClass:"add-QTY"},[t._v(t._s(t.QT3))]),t._v(" "),r("button",{staticClass:"bt_money",on:{click:function(e){t.add_QTY(2)}}},[r("img",{attrs:{src:n(182)}}),t._v(" "),r("h1",[t._v("เหรียญ 10 บาท")])]),t._v(" "),r("div",{staticClass:"rm-QTY",on:{click:function(e){t.rm_QTY(2)}}},[t._v(" - ")])]),t._v(" "),r("div",{staticClass:"data-money"},[r("div",{staticClass:"add-QTY"},[t._v(t._s(t.QT4))]),t._v(" "),r("button",{staticClass:"bt_money",on:{click:function(e){t.add_QTY(3)}}},[r("img",{staticStyle:{width:"80%",margin:"15% 0"},attrs:{src:n(271)}}),t._v(" "),r("h1",[t._v("ธนบัตร 100 บาท")])]),t._v(" "),r("div",{staticClass:"rm-QTY",on:{click:function(e){t.rm_QTY(3)}}},[t._v(" - ")])])])])])},staticRenderFns:[function(){var t=this,e=t.$createElement,n=t._self._c||e;return n("span",{staticClass:"icon icon is-large"},[n("i",{staticClass:"fa fa-arrow-left"})])}]}}});
//# sourceMappingURL=3.ee3ad4829dc2bf9a9e40.js.map