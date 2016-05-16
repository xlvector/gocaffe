import urllib2, urllib

imgs = []

for link in file("imgs.tsv"):
    imgs.append("http://image.guazistatic.com/" + link.strip())

print imgs

for i in range(100):
    buf = "|".join(imgs)
    postdata=urllib.urlencode({"imgs": buf})
    request = urllib2.Request("http://10.1.192.154:8011/predict",postdata)
    print urllib2.urlopen(request).read()
