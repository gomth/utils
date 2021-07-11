# Generic trivial utils collection

[![CircleCI](https://circleci.com/gh/gomth/utils/tree/main.svg?style=svg)](https://circleci.com/gh/gomth/utils/tree/main)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/715d6cdb76444f3999ba394513c6708d)](https://www.codacy.com/gh/gomth/utils/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=gomth/utils&amp;utm_campaign=Badge_Grade)
[![codebeat badge](https://codebeat.co/badges/80b9253b-9e7d-4d25-8a30-e09e4071be1c)](https://codebeat.co/projects/github-com-gomth-utils-main)
[![Coverage Status](https://coveralls.io/repos/github/gomth/utils/badge.svg?branch=main)](https://coveralls.io/github/gomth/utils?branch=main)

##### Package contents:
| API | Engine | time complexity | space complexity | Details |
| ------ | ------ | ------ | ------ | ------ |
| Gcd | [Euclidean algorythm][euclidian] | O(log\;min(a,b)) | O(1) | [Get greatest common divisor][gcd]. Recursive! |
| Powm | Modular exponentiation | - | - | [Memory-efficient modular exponentiation][powm]. Recursive! |

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [gcd]: <https://en.wikipedia.org/wiki/Greatest_common_divisor>
   [euclidian]: <https://en.wikipedia.org/wiki/Euclidean_algorithm>
   [powm]: <https://en.wikipedia.org/wiki/Modular_exponentiation>