#!/usr/bin/env zsh
# usage: get PARSER URL

PARSER=
URL=

if [ -n "$1" ]; then
	URL="$1"
	PARSER="${2:-html}"

	go run ./main.go -url "$1" | prettier --parser="$PARSER" | pygmentize -g

else
	echo "usage: $0 URL [PARSER]"
	echo "  URL - a completely formed url"
	echo "  PARSER - html, json, markdown, etc (default: html)"
	echo "  sample: json from http://httpbin.org"
	go run ./main.go | prettier --parser="json" | pygmentize
fi


 PARSERS=( "babel" "babel-flow" "babel-ts" "flow" "typescript" "css" "scss" "less" "json" "json5" "json-stringify" "graphql" "markdown" "mdx" "html" "vue" "angular" "lwc" "yaml" )

# Ref : Prettier parser choice ( use option --parser="xxxxx" option )
# Ref: prettier parser info: https://prettier.io/docs/en/options.html

    # Specify which parser to use.

    #     Prettier automatically infers the parser from the input file path, so you shouldn’t have to change this setting.

    #     Both the babel and flow parsers support the same set of JavaScript features (including Flow type annotations). They might differ in some edge cases, so if you run into one of those you can try flow instead of babel. Almost the same applies to typescript and babel-ts. babel-ts might support JavaScript features (proposals) not yet supported by TypeScript, but it’s less permissive when it comes to invalid code and less battle-tested than the typescript parser.

    # Valid options:

    #     "babel" (via @babel/parser) Named "babylon" until v1.16.0
    #     "babel-flow" (same as "babel" but enables Flow parsing explicitly to avoid ambiguity) First available in v1.16.0
    #     "babel-ts" (similar to "typescript" but uses Babel and its TypeScript plugin) First available in v2.0.0
    #     "flow" (via flow-parser)
    #     "typescript" (via @typescript-eslint/typescript-estree) First available in v1.4.0
    #     "css" (via postcss-scss and postcss-less, autodetects which to use) First available in v1.7.1
    #     "scss" (same parsers as "css", prefers postcss-scss) First available in v1.7.1
    #     "less" (same parsers as "css", prefers postcss-less) First available in v1.7.1
    #     "json" (via @babel/parser parseExpression) First available in v1.5.0
    #     "json5" (same parser as "json", but outputs as json5) First available in v1.13.0
    #     "json-stringify" (same parser as "json", but outputs like JSON.stringify) First available in v1.13.0
    #     "graphql" (via graphql/language) First available in v1.5.0
    #     "markdown" (via remark-parse) First available in v1.8.0
    #     "mdx" (via remark-parse and @mdx-js/mdx) First available in v1.15.0
    #     "html" (via angular-html-parser) First available in 1.15.0
    #     "vue" (same parser as "html", but also formats vue-specific syntax) First available in 1.10.0
    #     "angular" (same parser as "html", but also formats angular-specific syntax via angular-estree-parser) First available in 1.15.0
    #     "lwc" (same parser as "html", but also formats LWC-specific syntax for unquoted template attributes) First available in 1.17.0
    #     "yaml" (via yaml and yaml-unist-parser) First available in 1.14.0
