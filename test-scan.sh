#!/bin/bash

# Test script for vulnerable Go application
# This script demonstrates how to scan the application with Trivy

set -e

echo "========================================="
echo "Vulnerable Go Application - Test Script"
echo "========================================="
echo ""

# Check if trivy is installed
if ! command -v trivy &> /dev/null; then
    echo "❌ Trivy is not installed!"
    echo ""
    echo "To install Trivy:"
    echo "  macOS:   brew install trivy"
    echo "  Linux:   See https://aquasecurity.github.io/trivy/latest/getting-started/installation/"
    echo ""
    exit 1
fi

echo "✅ Trivy is installed"
echo ""

# Run basic scan
echo "1️⃣  Running basic vulnerability scan..."
echo "========================================="
trivy fs --severity HIGH,CRITICAL . || true
echo ""

# Run detailed scan with all severities
echo "2️⃣  Running detailed scan with all severity levels..."
echo "========================================="
trivy fs . || true
echo ""

# Generate JSON report
echo "3️⃣  Generating JSON report..."
echo "========================================="
trivy fs --format json --output trivy-results.json .
echo "✅ JSON report saved to: trivy-results.json"
echo ""

# Count vulnerabilities by severity
echo "4️⃣  Vulnerability Summary:"
echo "========================================="
if [ -f trivy-results.json ]; then
    echo "Analyzing results..."
    # This is a simple count - adjust based on your needs
    critical=$(grep -o '"Severity":"CRITICAL"' trivy-results.json 2>/dev/null | wc -l || echo "0")
    high=$(grep -o '"Severity":"HIGH"' trivy-results.json 2>/dev/null | wc -l || echo "0")
    medium=$(grep -o '"Severity":"MEDIUM"' trivy-results.json 2>/dev/null | wc -l || echo "0")
    low=$(grep -o '"Severity":"LOW"' trivy-results.json 2>/dev/null | wc -l || echo "0")
    
    echo "  CRITICAL: $critical"
    echo "  HIGH:     $high"
    echo "  MEDIUM:   $medium"
    echo "  LOW:      $low"
fi
echo ""

echo "========================================="
echo "Scan Complete!"
echo "========================================="
echo ""
echo "Next steps for your AI agent:"
echo "  1. Parse trivy-results.json"
echo "  2. Identify vulnerable dependencies in go.mod"
echo "  3. Update to secure versions"
echo "  4. Test that the application still builds"
echo "  5. Re-run Trivy to verify fixes"
echo ""
