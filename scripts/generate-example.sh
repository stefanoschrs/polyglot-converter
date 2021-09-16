#!/bin/bash

set -e

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"
cd ${DIR}/..

rm example/example.{go,ts,py}

go run main.go example/example.yml example/

readmeFile="example/README.md"
echo "# Full Example" > "${readmeFile}"

echo "## YAML" >> "${readmeFile}"
echo '```yml' >> "${readmeFile}"
cat example/example.yml >> "${readmeFile}"
echo '```' >> "${readmeFile}"

echo "## Go" >> "${readmeFile}"
echo '```go' >> "${readmeFile}"
cat example/example.go >> "${readmeFile}"
echo '```' >> "${readmeFile}"

echo "## TypeScript" >> "${readmeFile}"
echo '```typescript' >> "${readmeFile}"
cat example/example.ts >> "${readmeFile}"
echo '```' >> "${readmeFile}"

echo "## Python" >> "${readmeFile}"
echo '```python3' >> "${readmeFile}"
cat example/example.py >> "${readmeFile}"
echo '```' >> "${readmeFile}"
