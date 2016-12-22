package main

const appHtml = `
<html >

<head>

 <title>Customer Service Portal</title>
<script type="text/javascript" src="https://code.jquery.com/jquery-2.1.3.min.js"></script>
<script type="text/javascript" src="/static/js/test5.js"></script>

<link href="/static/css/vendor/bootstrap.min.css" rel="stylesheet">

<style>
body {
	
	margin:0px 200px;
	
	background-image: url("/static/js/bg.jpg");
    background-color: #cccccc;
	 
	}
.progress {
  height: 12px;
  background: #ebedef;
  border-radius: 32px;
  box-shadow: none;
  }
.progress-bar {
  line-height: 12px;
  background: #1abc9c;
  box-shadow: none;
  }
.progress-bar-success {
  background-color: #2ecc71;
  }
.progress-bar-warning {
  background-color: #f1c40f;
  }
.progress-bar-danger {
  background-color: #e74c3c;
  }
.progress-bar-info {
  background-color: #3498db;
  }
</style>


</head>

<body  >
  <div id="main">

	<div width: 960px;  height: 50px;  margin-top: 50px; >
	  <div width: 960px;  height: 50px;  margin-top: 50px; align='middle'>
	    <h1><a href="#">Customer Service Portal</a></h1>
	  </div><!--close welcome-->
     
    </div><!--close menubar-->	
    
	   <div id="content" >
        <div class="content_item" align='middle'><br><br>
		  <h1>Welcome To EFI's Customer Service Portal</h1> 
	      <p>Browse for the file to upload :</p>	  
		  <div class="content_container" align='10'>

		
<form action="/upload/" method="POST" enctype="multipart/form-data"  id="id-filedata" name="id-filedata">
						<input type="file" name="id-file-d" id="id-file-d"><br>
						
					</form>
					
    	    <!--close button_small-->
		  </div><!--close content_container-->
    	  
		</div><!--close content_item-->
    </div><!--close main-->
	
	
	<p>
					<div class="progress" id="id-progress-bar" align="center"  style="display:none;">
		            		<div class="progress-bar" id="id-upload-progress" align="center" color="black" style="">
							
							</div>
		             </div>
					<div id="id-percentile" align="center"></div>
					</p>
  
 
  
</body>
</html>
`
