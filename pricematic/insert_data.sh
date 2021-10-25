#!/usr/bin/env bash
wget data.json
echo '\set content `cat $(ls -t algo.json.* | head -1)`'