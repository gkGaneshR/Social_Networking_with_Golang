package frontend

const IndexPage = `
<html >
	<head>
		<title>Customer Service Portal</title>
		<style>
		body {
			margin:0px ;
			background-image: url("/static/css/bg.jpg");
		    background-color: #cccccc;
			text-align: center;
			}
		#efi {
			color : white ;
			}
		.toptag {
			margin-top : 0px ;
			margin-left : 0px ;
			margin-right : 0px ;
			height : 90px ;
			background-color : black ;
			opacity : 0.95;
			text-align:left;
			}
		</style>
	</head>
	<body>
			<div class="toptag">
						<img src="/static/js/efi logo.png" style="float:left;" />
						<h1 id="efi"><br>Customer Service Portal</h1>
			</div>
	
		   <div id="content" >
		        <div class="content_item" align='middle'><br><br>
				    <h1>Welcome To EFI's Customer Service Portal</h1> 
						<br><br>
					<h1>%s</h1><br>
						<form method="post" action="/login">
						    <label for="name">User name</label>
						    <input type="text" id="name" name="name">
						    <label for="password">Password</label>
						    <input type="password" id="password" name="password">
						    <button type="submit">Login</button>
						</form><br>
					<a href="/register/">New User? Register here</a>
				</div>
			</div>
		<style>
						
					html, body, div, span, applet, object, iframe,
					h1, h2, h3, h4, h5, h6, p, blockquote, pre,
					a, abbr, acronym, address, big, cite, code,
					del, dfn, em, img, ins, kbd, q, s, samp,
					small, strike, strong, sub, sup, tt, var,
					b, u, i, center,
					dl, dt, dd, ol, ul, li,
					fieldset, form, label, legend,
					table, caption, tbody, tfoot, thead, tr, th, td,
					article, aside, canvas, details, embed, 
					figure, figcaption, footer, header, hgroup, 
					menu, nav, output, ruby, section, summary,
					time, mark, audio, video {
						margin: 0;
						padding: 0;
						border: 0;
						font-size: 100 %;
						font: inherit;
						vertical-align: baseline;
					}
					/* HTML5 display-role reset for older browsers */
					article, aside, details, figcaption, figure, 
					footer, header, hgroup, menu, nav, section {
						display: block;
					}
					body {
						line-height: 1;
					}
					ol, ul {
						list-style: none;
					}
					blockquote, q {
						quotes: none;
					}
					blockquote:before, blockquote:after,
					q:before, q:after {
						content: '';
						content: none;
					}
					table {
						border-collapse: collapse;
						border-spacing: 0;
					}
		</style>
	</body>
</html>
`
