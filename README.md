jsp-timestamper
====
[![Gobuild Download](http://gobuild.io/badge/github.com/mallowlabs/jsp-timestamper/download.png)](http://gobuild.io/github.com/mallowlabs/jsp-timestamper)

A tool that appends timestamp parameters to script and style tags in JSP files.


## Description
The browser caching sometimes raises many problem with JavaScripts and style sheets.
Some web application frameworks appends timestamps to these files.
But if your Java web application does not support this mechanism,
you will be annoyed with caching.
This tool provides timestamp appending to deployed JSP files.


For example:
```html
<!-- before with tool -->
<link type="text/css" href="css/ui-lightness/jquery-ui-1.8.18.custom.css" rel="stylesheet" />
<script type="text/javascript" src="js/jquery-1.7.1.min.js"></script>
<script type="text/javascript" src="js/jquery-ui-1.8.18.custom.min.js"></script>
```
```html
<!-- after with tool -->
<link type="text/css" href="css/ui-lightness/jquery-ui-1.8.18.custom.css?20140914093459" rel="stylesheet" />
<script type="text/javascript" src="js/jquery-1.7.1.min.js?20140914093459"></script>
<script type="text/javascript" src="js/jquery-ui-1.8.18.custom.min.js?20140914093459"></script>
```

## VS.
* [Assset Pipeline](http://guides.rubyonrails.org/asset_pipeline.html)

## Usage
1. Deploy your war file.
2. Run this tool to deployed JSP files.

```
./jsp-timestamper <JSP_DIRECTORY_PATH>
```
For example, JSP_DIRECTORY_PATH: ```/usr/local/tomcat/webapps/your-app/WEB-INF/page```

IMPORTANT: You must run this tool BEFORE any web accesses. (i.e. before JSP pre-compile)

## Licence

[ MIT Licence](https://github.com/mallowlabs/jsp-timestamper/blob/master/LICENCE.txt)

## Author

[mallowlabs](https://github.com/mallowlabs)
