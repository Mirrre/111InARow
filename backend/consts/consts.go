package consts

import "time"

const AccessKey = "UbG9NVTSj7i4kJ_BAUWcaQi6CrVPgOPRh5eWrDdA"
const SecretKey = "UxuMXKJ4Sx7gAPmBht4mzEKWaK9D9_MNGXQpy8lo"
const CacheKeyFollow = "user:%d:follow"
const CacheKeyFollower = "user:%d:follower"
const CacheKeyComments = "video:%d:comments"
const MaxCommentLength = 512
const CacheExpiration = 5 * time.Minute
