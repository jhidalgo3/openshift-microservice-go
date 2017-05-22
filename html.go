/**
# Copyright 2015 Google Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
**/

package main

const (
	html = `<!doctype html>
<html>
<head>
<!-- Compiled and minified CSS -->
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.97.0/css/materialize.min.css">

<!-- Compiled and minified JavaScript -->
<script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/0.97.0/js/materialize.min.js"></script>
<title>Frontend Web Server</title>
</head>
<body>
<div class="container">
<div class="row">
<div class="col s2">&nbsp;</div>
<div class="col s8">


<div class="card green">
<div class="card-content white-text">
<div class="card-title">Pod that serviced this request</div>
</div>
<div class="card-content white">
<table class="bordered">
	<tbody>
	<tr>
		<td>App</td>
		<td>{{.App}}</td>
	</tr>
	<tr>
		<td>Namespace</td>
		<td>{{.Namespace}}</td>
	</tr>
	<tr>
		<td>Name</td>
		<td>{{.Name}}</td>
	</tr>
	<tr>
		<td>Version</td>
		<td>{{.Version}}</td>
	</tr>
	<tr>
		<td>Hostname</td>
		<td>{{.Hostname}}</td>
	</tr>
	<tr>
		<td>IP</td>
		<td>{{.Ip}}</td>
	</tr>
	

</table>
</div>
</div>
</html>`
)
