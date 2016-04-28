import urllib2, urllib

imgs = [
    "http://image.guazistatic.com/gz01160427/15/00/c19effea5e7ca9655864119dbd657eed.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/fcd0e61ab4ee96f20480a06c4286aaa7.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/30b29858488e9e0b53bb31a613ff7cb1.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/a3ca944f3329768b4004257c072aa5dc.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/fed78ddd0eefe44577aea57eedd634f0.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/72f93eac60b23d71221e793939b45c0d.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/2ddce68b2e1e3f00d5c2bf3ebc4b0eba.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/266fa16cf831044338ad6cde52aa0c8f.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/35f467323a6e5075237749bf9177d10d.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/8a5aa98fb6fda482af01393e70af612d.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/64a3e548dc79e98cf9b98ac824ec36f7.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/a4fbf6f49f0e7ed6a125665a10cb465d.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/26b527a09402aa7834cd23d65ec4fb8b.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/53e4a0f917f8ff30f470d4dd6b9662a8.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/9a1aab18904ef927e02de1a85e4764c6.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/d174864df0cbd8311acdb985bd219c72.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/2a9d1cd256a5248dba83dada9e26558f.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/fd63f432ec2194515b3b8fbfebc7777f.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/64da27886729dcc895cb43c7aabf8924.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/dd6c96e7e07b42452e3b3e90f7634a44.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/3bd5cd94257b5ec0503d6a6f25ce4d68.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/69a51b6cc8b2fe3e7cecced9f9c3daa6.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/5889327e4d38337a58c8ffbd54d4bbe3.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/0a250350309f64b0433d2a52354602d3.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/4c3079d84be7860b6d63eff48b60bf66.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/ef09eacaf15b9d485929c797d4e0d0d8.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/0f31e512eba375ef51016494246575b9.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/69675ca6bb1fe99c2bc8319a0501fcd0.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/44755aec8c67cfa77a017cb685ce59b1.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88",
    "http://image.guazistatic.com/gz01160427/15/00/74b46acd0c88fd5a030bdea41f303e70.jpg@base@tag=imgScale&w=150&h=100&c=1&m=2&q=88"
]

buf = "|".join(imgs)
postdata=urllib.urlencode({"imgs": buf})
request = urllib2.Request("http://10.1.192.128:8011/predict",postdata)
print urllib2.urlopen(request).read()
