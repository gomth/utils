# Generic trivial utils collection

[![CircleCI](https://circleci.com/gh/gomth/utils/tree/main.svg?style=svg)](https://circleci.com/gh/gomth/utils/tree/main)
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