package main

// MovieHTMLHead html header
func MovieHTMLHead() string {
	head := `
<!DOCTYPE html> 
<html>
<head>
    <title>New Movies Uploaded</title>
</head>
<body>
    <center>
        <h1>New Movies Uploaded</h1>
    </center>
`
	return head
}

// MovieHTMLTable html body
func MovieHTMLTable() string {
	table := `
<table >
	<tbody>
		<tr>
			<td><img src=%s alt="" border=3 height=auto width=100></img></td>
			<td>
				<table>
					<tbody>
						<tr>
							<td>Name: %s</td>
						</tr>
						<tr>
							<td>Year: %d</td>
						</tr>
						<tr>
							<td>Genres : %s</td>
						</tr>
						<tr>
							<td>IMDB Rating: %g, IMDB Id: %s</td>
						</tr>
						<tr>
							<td>Upload Date: %s</td>
						</tr>
					</tbody>
				</table>
			</td>
		</tr>
	</tbody>
</table>

`
	return table

}

// HTMLEnd close body and html
func HTMLEnd() string {
	end := "</body> </html>"
	return end
}
