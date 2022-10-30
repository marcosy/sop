#!/bin/bash
set -o pipefail
set -o errexit 
set -o nounset

run-test() {
    sop=../../bin/sop

    # Execute
    $sop > /dev/null
    $sop $operation $file1 $file2 | sort > $actual 

    # Assert
    set +o errexit 
    diff $expected $actual
    local status=$?
    set -o errexit

    if [ $status -eq 0 ]; then 
        rm $actual 
        exit 0
    fi

    diff --side-by-side $expected $actual
    exit $status
}

# Run scenarios
echo "Running tests scenarios..."
scenarios=($( ls test-* ))
total_scenarios=${#scenarios[@]}
current_scenario=1
for scenario in ${scenarios[*]}; do
    echo "Running [$current_scenario / $total_scenarios] $scenario "
    source "${scenario}"

    set +o errexit
    run-test
    s=$?
    set -o errexit
    
    if [ $s -ne 0 ]; then
        failed+=( $scenario )
    fi
    ((current_scenario++))
done
echo "Done!"
echo ""

# Show Results
if [ ${#failed[@]} -ne 0 ]; then
    echo "Test failed: ${failed[*]}"
    exit 1
fi
echo "Test Passed"



