/*! modernizr 3.4.0 (Custom Build) | MIT *
 * https://modernizr.com/download/?-appearance-backgroundsize-canvas-cookies-cssanimations-csscalc-cssfilters-cssgradients-csstransforms-csstransitions-eventlistener-forcetouch-fullscreen-geolocation-hiddenscroll-input-inputtypes-json-ligatures-lowbattery-notification-passiveeventlisteners-pointerevents-postmessage-queryselector-search-serviceworker-svg-touchevents-video-setclasses !*/
!function(e,t,n){function r(e,t){return typeof e===t}function i(){var e,t,n,i,o,s,a;for(var l in S)if(S.hasOwnProperty(l)){if(e=[],t=S[l],t.name&&(e.push(t.name.toLowerCase()),t.options&&t.options.aliases&&t.options.aliases.length))for(n=0;n<t.options.aliases.length;n++)e.push(t.options.aliases[n].toLowerCase());for(i=r(t.fn,"function")?t.fn():t.fn,o=0;o<e.length;o++)s=e[o],a=s.split("."),1===a.length?Modernizr[a[0]]=i:(!Modernizr[a[0]]||Modernizr[a[0]]instanceof Boolean||(Modernizr[a[0]]=new Boolean(Modernizr[a[0]])),Modernizr[a[0]][a[1]]=i),T.push((i?"":"no-")+a.join("-"))}}function o(e){var t=b.className,n=Modernizr._config.classPrefix||"";if(x&&(t=t.baseVal),Modernizr._config.enableJSClass){var r=new RegExp("(^|\\s)"+n+"no-js(\\s|$)");t=t.replace(r,"$1"+n+"js$2")}Modernizr._config.enableClasses&&(t+=" "+n+e.join(" "+n),x?b.className.baseVal=t:b.className=t)}function s(){return"function"!=typeof t.createElement?t.createElement(arguments[0]):x?t.createElementNS.call(t,"http://www.w3.org/2000/svg",arguments[0]):t.createElement.apply(t,arguments)}function a(e){return e.replace(/([a-z])-([a-z])/g,function(e,t,n){return t+n.toUpperCase()}).replace(/^-/,"")}function l(){var e=t.body;return e||(e=s(x?"svg":"body"),e.fake=!0),e}function u(e,n,r,i){var o,a,u,c,d="modernizr",f=s("div"),p=l();if(parseInt(r,10))for(;r--;)u=s("div"),u.id=i?i[r]:d+(r+1),f.appendChild(u);return o=s("style"),o.type="text/css",o.id="s"+d,(p.fake?p:f).appendChild(o),p.appendChild(f),o.styleSheet?o.styleSheet.cssText=e:o.appendChild(t.createTextNode(e)),f.id=d,p.fake&&(p.style.background="",p.style.overflow="hidden",c=b.style.overflow,b.style.overflow="hidden",b.appendChild(p)),a=n(f,e),p.fake?(p.parentNode.removeChild(p),b.style.overflow=c,b.offsetHeight):f.parentNode.removeChild(f),!!a}function c(e,t){return!!~(""+e).indexOf(t)}function d(e,t){return function(){return e.apply(t,arguments)}}function f(e,t,n){var i;for(var o in e)if(e[o]in t)return n===!1?e[o]:(i=t[e[o]],r(i,"function")?d(i,n||t):i);return!1}function p(e){return e.replace(/([A-Z])/g,function(e,t){return"-"+t.toLowerCase()}).replace(/^ms-/,"-ms-")}function v(t,n,r){var i;if("getComputedStyle"in e){i=getComputedStyle.call(e,t,n);var o=e.console;if(null!==i)r&&(i=i.getPropertyValue(r));else if(o){var s=o.error?"error":"log";o[s].call(o,"getComputedStyle returning null, its possible modernizr test results are inaccurate")}}else i=!n&&t.currentStyle&&t.currentStyle[r];return i}function m(t,r){var i=t.length;if("CSS"in e&&"supports"in e.CSS){for(;i--;)if(e.CSS.supports(p(t[i]),r))return!0;return!1}if("CSSSupportsRule"in e){for(var o=[];i--;)o.push("("+p(t[i])+":"+r+")");return o=o.join(" or "),u("@supports ("+o+") { #modernizr { position: absolute; } }",function(e){return"absolute"==v(e,null,"position")})}return n}function g(e,t,i,o){function l(){d&&(delete W.style,delete W.modElem)}if(o=r(o,"undefined")?!1:o,!r(i,"undefined")){var u=m(e,i);if(!r(u,"undefined"))return u}for(var d,f,p,v,g,h=["modernizr","tspan","samp"];!W.style&&h.length;)d=!0,W.modElem=s(h.shift()),W.style=W.modElem.style;for(p=e.length,f=0;p>f;f++)if(v=e[f],g=W.style[v],c(v,"-")&&(v=a(v)),W.style[v]!==n){if(o||r(i,"undefined"))return l(),"pfx"==t?v:!0;try{W.style[v]=i}catch(y){}if(W.style[v]!=g)return l(),"pfx"==t?v:!0}return l(),!1}function h(e,t,n,i,o){var s=e.charAt(0).toUpperCase()+e.slice(1),a=(e+" "+M.join(s+" ")+s).split(" ");return r(t,"string")||r(t,"undefined")?g(a,t,i,o):(a=(e+" "+L.join(s+" ")+s).split(" "),f(a,t,n))}function y(e,t,r){return h(e,n,n,t,r)}var T=[],S=[],w={_version:"3.4.0",_config:{classPrefix:"",enableClasses:!0,enableJSClass:!0,usePrefixes:!0},_q:[],on:function(e,t){var n=this;setTimeout(function(){t(n[e])},0)},addTest:function(e,t,n){S.push({name:e,fn:t,options:n})},addAsyncTest:function(e){S.push({name:null,fn:e})}},Modernizr=function(){};Modernizr.prototype=w,Modernizr=new Modernizr,Modernizr.addTest("cookies",function(){try{t.cookie="cookietest=1";var e=-1!=t.cookie.indexOf("cookietest=");return t.cookie="cookietest=1; expires=Thu, 01-Jan-1970 00:00:01 GMT",e}catch(n){return!1}}),Modernizr.addTest("eventlistener","addEventListener"in e),Modernizr.addTest("geolocation","geolocation"in navigator),Modernizr.addTest("json","JSON"in e&&"parse"in JSON&&"stringify"in JSON),Modernizr.addTest("notification",function(){if(!e.Notification||!e.Notification.requestPermission)return!1;if("granted"===e.Notification.permission)return!0;try{new e.Notification("")}catch(t){if("TypeError"===t.name)return!1}return!0}),Modernizr.addTest("postmessage","postMessage"in e),Modernizr.addTest("queryselector","querySelector"in t&&"querySelectorAll"in t),Modernizr.addTest("serviceworker","serviceWorker"in navigator),Modernizr.addTest("svg",!!t.createElementNS&&!!t.createElementNS("http://www.w3.org/2000/svg","svg").createSVGRect),Modernizr.addTest("passiveeventlisteners",function(){var t=!1;try{var n=Object.defineProperty({},"passive",{get:function(){t=!0}});e.addEventListener("test",null,n)}catch(r){}return t});var b=t.documentElement,x="svg"===b.nodeName.toLowerCase();Modernizr.addTest("canvas",function(){var e=s("canvas");return!(!e.getContext||!e.getContext("2d"))}),Modernizr.addTest("video",function(){var e=s("video"),t=!1;try{t=!!e.canPlayType,t&&(t=new Boolean(t),t.ogg=e.canPlayType('video/ogg; codecs="theora"').replace(/^no$/,""),t.h264=e.canPlayType('video/mp4; codecs="avc1.42E01E"').replace(/^no$/,""),t.webm=e.canPlayType('video/webm; codecs="vp8, vorbis"').replace(/^no$/,""),t.vp9=e.canPlayType('video/webm; codecs="vp9"').replace(/^no$/,""),t.hls=e.canPlayType('application/x-mpegURL; codecs="avc1.42E01E"').replace(/^no$/,""))}catch(n){}return t});var C=function(){function e(e,t){var i;return e?(t&&"string"!=typeof t||(t=s(t||"div")),e="on"+e,i=e in t,!i&&r&&(t.setAttribute||(t=s("div")),t.setAttribute(e,""),i="function"==typeof t[e],t[e]!==n&&(t[e]=n),t.removeAttribute(e)),i):!1}var r=!("onblur"in t.documentElement);return e}();w.hasEvent=C,Modernizr.addTest("inputsearchevent",C("search"));var E=s("input"),_="autocomplete autofocus list placeholder max min multiple pattern required step".split(" "),k={};Modernizr.input=function(t){for(var n=0,r=t.length;r>n;n++)k[t[n]]=!!(t[n]in E);return k.list&&(k.list=!(!s("datalist")||!e.HTMLDataListElement)),k}(_);var N="search tel url email datetime date month week time datetime-local number range color".split(" "),P={};Modernizr.inputtypes=function(e){for(var r,i,o,s=e.length,a="1)",l=0;s>l;l++)E.setAttribute("type",r=e[l]),o="text"!==E.type&&"style"in E,o&&(E.value=a,E.style.cssText="position:absolute;visibility:hidden;",/^range$/.test(r)&&E.style.WebkitAppearance!==n?(b.appendChild(E),i=t.defaultView,o=i.getComputedStyle&&"textfield"!==i.getComputedStyle(E,null).WebkitAppearance&&0!==E.offsetHeight,b.removeChild(E)):/^(search|tel)$/.test(r)||(o=/^(url|email)$/.test(r)?E.checkValidity&&E.checkValidity()===!1:E.value!=a)),P[e[l]]=!!o;return P}(N);var O=w._config.usePrefixes?" -webkit- -moz- -o- -ms- ".split(" "):["",""];w._prefixes=O,Modernizr.addTest("csscalc",function(){var e="width:",t="calc(10px);",n=s("a");return n.style.cssText=e+O.join(t+e),!!n.style.length}),Modernizr.addTest("cssgradients",function(){for(var e,t="background-image:",n="gradient(linear,left top,right bottom,from(#9f9),to(white));",r="",i=0,o=O.length-1;o>i;i++)e=0===i?"to ":"",r+=t+O[i]+"linear-gradient("+e+"left top, #9f9, white);";Modernizr._config.usePrefixes&&(r+=t+"-webkit-"+n);var a=s("a"),l=a.style;return l.cssText=r,(""+l.backgroundImage).indexOf("gradient")>-1});var A="CSS"in e&&"supports"in e.CSS,z="supportsCSS"in e;Modernizr.addTest("supports",A||z);var j="Moz O ms Webkit",L=w._config.usePrefixes?j.toLowerCase().split(" "):[];w._domPrefixes=L,Modernizr.addTest("pointerevents",function(){var e=!1,t=L.length;for(e=Modernizr.hasEvent("pointerdown");t--&&!e;)C(L[t]+"pointerdown")&&(e=!0);return e});var M=w._config.usePrefixes?j.split(" "):[];w._cssomPrefixes=M;var R=function(t){var r,i=O.length,o=e.CSSRule;if("undefined"==typeof o)return n;if(!t)return!1;if(t=t.replace(/^@/,""),r=t.replace(/-/g,"_").toUpperCase()+"_RULE",r in o)return"@"+t;for(var s=0;i>s;s++){var a=O[s],l=a.toUpperCase()+"_"+r;if(l in o)return"@-"+a.toLowerCase()+"-"+t}return!1};w.atRule=R;var $=w.testStyles=u;Modernizr.addTest("hiddenscroll",function(){return $("#modernizr {width:100px;height:100px;overflow:scroll}",function(e){return e.offsetWidth===e.clientWidth})}),Modernizr.addTest("touchevents",function(){var n;if("ontouchstart"in e||e.DocumentTouch&&t instanceof DocumentTouch)n=!0;else{var r=["@media (",O.join("touch-enabled),("),"heartz",")","{#modernizr{top:9px;position:absolute}}"].join("");$(r,function(e){n=9===e.offsetTop})}return n});var q={elem:s("modernizr")};Modernizr._q.push(function(){delete q.elem});var W={style:q.elem.style};Modernizr._q.unshift(function(){delete W.style}),w.testAllProps=h;var U=w.prefixed=function(e,t,n){return 0===e.indexOf("@")?R(e):(-1!=e.indexOf("-")&&(e=a(e)),t?h(e,t,n):h(e,"pfx"))};Modernizr.addTest("forcetouch",function(){return C(U("mouseforcewillbegin",e,!1),e)?MouseEvent.WEBKIT_FORCE_AT_MOUSE_DOWN&&MouseEvent.WEBKIT_FORCE_AT_FORCE_MOUSE_DOWN:!1}),Modernizr.addTest("fullscreen",!(!U("exitFullscreen",t,!1)&&!U("cancelFullScreen",t,!1))),Modernizr.addTest("lowbattery",function(){var e=.2,t=U("battery",navigator);return!!(t&&!t.charging&&t.level<=e)}),w.testAllProps=y,Modernizr.addTest("ligatures",y("fontFeatureSettings",'"liga" 1')),Modernizr.addTest("cssanimations",y("animationName","a",!0)),Modernizr.addTest("appearance",y("appearance")),Modernizr.addTest("backgroundsize",y("backgroundSize","100%",!0)),Modernizr.addTest("cssfilters",function(){if(Modernizr.supports)return y("filter","blur(2px)");var e=s("a");return e.style.cssText=O.join("filter:blur(2px); "),!!e.style.length&&(t.documentMode===n||t.documentMode>9)}),Modernizr.addTest("csstransforms",function(){return-1===navigator.userAgent.indexOf("Android 2.")&&y("transform","scale(1)",!0)}),Modernizr.addTest("csstransitions",y("transition","all",!0)),i(),o(T),delete w.addTest,delete w.addAsyncTest;for(var V=0;V<Modernizr._q.length;V++)Modernizr._q[V]();e.Modernizr=Modernizr}(window,document);