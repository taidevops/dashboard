import gitDescribe from 'git-describe';

import {writeFileSync, readFileSync} from 'node:fs';

const config = JSON.parse(readFileSync('package.json', 'utf8'));

// Another example: working directory, use 16 character commit hash abbreviation
const gitInfo = gitDescribe.gitDescribeSync({
  dirtyMark: false,
  dirtySemver: false,
  longSemver: true,
});

gitInfo.packageVersion = config.version;
console.log(gitInfo);