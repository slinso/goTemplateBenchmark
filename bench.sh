#!/usr/bin/env bash
# This file:
#
#  - Run the benchmarks and output github compatible markdown for the readme.md
#
# Usage:
#
#  ./bench.sh
#
# Testing:
# only run the benchmarks
#  ./bench.sh -O -F
#
# Based on a template by BASH3 Boilerplate v2.3.0
# http://bash3boilerplate.sh/#authors
#
# The MIT License (MIT)
# Copyright (c) 2013 Kevin van Zonneveld and contributors
# You are not obligated to bundle the LICENSE file with your b3bp projects as long
# as you leave these references intact in the header comments of your source files.

# shellcheck disable=SC2034
read -r -d '' __usage <<-'EOF' || true # exits non-zero when EOF encountered
  -t --time  [arg]   Benchmark duration. Required. Default="3s"
  -c --compare [arg] Old go version binary? Required.
  -B --no-benchmarks Do NOT run the benchmarks.
  -F --no-format     Do NOT format the results.
  -v                 Enable verbose mode, print script as it is executed
  -d --debug         Enables debug mode
  -h --help          This page
EOF

# shellcheck disable=SC2034
read -r -d '' __helptext <<-'EOF' || true # exits non-zero when EOF encountered
 This is Bash3 Boilerplate's help text. Feel free to add any description of your
 program or elaborate more on command-line arguments. This section is not
 parsed and will be added as-is to the help.
EOF

# shellcheck source=main.sh
source "$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/main.sh"

### Signal trapping and backtracing
##############################################################################

# requires `set -o errtrace`
__b3bp_err_report() {
    local error_code=${?}
    # shellcheck disable=SC2154
    error "Error in ${__file} in function ${1} on line ${2}"
    exit ${error_code}
}
# Uncomment the following line for always providing an error backtrace
trap '__b3bp_err_report "${FUNCNAME:-.}" ${LINENO}' ERR

### Command-line argument switches (like -d for debugmode, -h for showing helppage)
##############################################################################

# debug mode
if [[ "${arg_d:?}" == "1" ]]; then
    set -o xtrace
    PS4='+(${BASH_SOURCE}:${LINENO}): ${FUNCNAME[0]:+${FUNCNAME[0]}(): }'
    LOG_LEVEL="7"
    # Enable error backtracing
    trap '__b3bp_err_report "${FUNCNAME:-.}" ${LINENO}' ERR
fi

# verbose mode
if [[ "${arg_v:?}" == "1" ]]; then
    set -o verbose
fi

# help mode
if [[ "${arg_h:?}" == "1" ]]; then
    # Help exists with code 1
    help "Help using ${0}"
fi

__no_benchmarks="false"
if [[ "${arg_B:?}" == "1" ]]; then
    __no_benchmarks="true"
fi

__no_format="false"
if [[ "${arg_F:?}" == "1" ]]; then
    __no_format="true"
fi

### Validation. Error out if the things required for your script are not present
##############################################################################

[[ "${arg_t:-}" ]] || help "Setting benchmark druation with -t or --time is required"
[[ "${arg_c:-}" ]] || help "Setting go versions which will be compared with -c or --compare is required"
[[ "${LOG_LEVEL:-}" ]] || emergency "Cannot continue without LOG_LEVEL. "

### Runtime
##############################################################################

info "OSTYPE: ${OSTYPE}"

info "benchmark duration: ${arg_t}"
info "compare: ${arg_c} to $(go version)"
info "run benchmarks: $([[ "${__no_benchmarks}" == "true" ]] && echo "false" || echo "true")"
info "format output: $([[ "${__no_format}" == "true" ]] && echo "false" || echo "true")"

# run old benchmarks
_run_old_benchmarks() {
    ${arg_c} test -bench "k(Ace|Amber|Golang|GolangText|Handlebars|Mustache|Pongo2|Soy|JetHTML)$" -benchmem -benchtime="${arg_t}" | tee ./files/results-1.old
    ${arg_c} test -bench "k(Ego|EgonSlinso|Quicktemplate|Ftmpl|Gorazor|Hero|Jade)$" -benchmem -benchtime="${arg_t}" | tee ./files/results-2.old
    ${arg_c} test -bench "Complex(Ace|Amber|Golang|GolangText|Handlebars|Mustache|Pongo2|Soy|JetHTML)$" -benchmem -benchtime="${arg_t}" | tee ./files/results-3.old
    ${arg_c} test -bench "Complex(Ego|EgoSlinso|Quicktemplate|Ftmpl|Gorazor|Hero|Jade)$" -benchmem -benchtime="${arg_t}" | tee ./files/results-4.old
    ${arg_c} test -bench "Complex(GoStaticString|GoDirectBuffer)$" -benchmem -benchtime="${arg_t}" | tee ./files/results-5.old
}
[[ "${__no_benchmarks}" == "true" ]] || _run_old_benchmarks

# run benchmarks
_run_benchmarks() {
    go test -bench "k(Ace|Amber|Golang|GolangText|Handlebars|Mustache|Pongo2|Soy|JetHTML)$" -benchmem -benchtime="${arg_t}" | tee ./files/results-1.new
    go test -bench "k(Ego|EgonSlinso|Quicktemplate|Ftmpl|Gorazor|Hero|Jade)$" -benchmem -benchtime="${arg_t}" | tee ./files/results-2.new
    go test -bench "Complex(Ace|Amber|Golang|GolangText|Handlebars|Mustache|Pongo2|Soy|JetHTML)$" -benchmem -benchtime="${arg_t}" | tee ./files/results-3.new
    go test -bench "Complex(Ego|EgoSlinso|Quicktemplate|Ftmpl|Gorazor|Hero|Jade)$" -benchmem -benchtime="${arg_t}" | tee ./files/results-4.new
    go test -bench "Complex(GoStaticString|GoDirectBuffer)$" -benchmem -benchtime="${arg_t}" | tee ./files/results-5.new
}
[[ "${__no_benchmarks}" == "true" ]] || _run_benchmarks

# formats a single benchmark
# $1: number of the benchmark
__format_single_benchmark() {
    local i=${1}
    pb <./files/results-"${i}".new | grep \| | sed '/Name/a \| --- \| --- \| --- \| --- \| --- \|'
    echo ""
    echo "\`\`\`"
    echo "comparing: ${arg_c} to $(go version)"
    benchstat -delta-test none files/results-"${i}".old files/results-"${i}".new | tee files/results-"${i}"-benchstat.txt
    echo "\`\`\`"
}

# pretty print for readme.md
_format_benchmarks() {
    echo ""
    echo "## special benchmarks"
    __format_single_benchmark 5

    echo ""
    echo "## simple benchmarks"
    echo "### full featured template engines"
    __format_single_benchmark 1

    echo ""
    echo "### precompilation to Go code"
    __format_single_benchmark 2

    echo ""
    echo "## more complex test with template inheritance (if possible)"
    echo "### full featured template engines"
    __format_single_benchmark 3

    echo ""
    echo "### precompilation to Go code"
    __format_single_benchmark 4
}
[[ "${__no_format}" == "true" ]] || _format_benchmarks
