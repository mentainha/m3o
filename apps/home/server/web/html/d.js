var maxChars = 1024;

String.prototype.parseURL = function(embed) {
	return this.replace(/[A-Za-z]+:\/\/[A-Za-z0-9-_]+\.[A-Za-z0-9-_:%&~\?\/.=]+/g, function(url) {
		if (embed == true) {
			var match = url.match(/^.*(youtu.be\/|v\/|u\/\w\/|embed\/|watch\?v=|\&v=)([^#\&\?]*).*/);
			if (match && match[2].length == 11) {
				return '<div class="iframe">'+
				'<iframe src="//www.youtube.com/embed/' + match[2] +
				'" frameborder="0" allowfullscreen></iframe>' + '</div>';
			};
			if (url.match(/^.*giphy.com\/media\/[a-zA-Z0-9]+\/[a-zA-Z0-9]+.gif$/)) {
				return '<div class="animation"><img src="'+url+'"></div>';
			}
			if (url.match(/https:\/\/.*.(jpg|jpeg|png)/)) {
				return '<div class="image"><a href="'+url+'"><img src="'+url+'"></a></div>';
			}
		};
		var pretty = url.replace(/^http(s)?:\/\/(www\.)?/, '');
		return pretty.link(url);
	});
};

String.prototype.parseHashTag = function() {
	return this.replace(/#[A-Za-z0-9-_+]+/g, function(t) {
		//var tag = t.replace("#","%23")
		var url = location.protocol+'//'+location.hostname+(location.port ? ':'+location.port: '');
		return t.link(url + '/' + t);
	});
};

String.prototype.capitalize = function () {
    return this.charAt(0).toUpperCase() + this.slice(1);
};

// load the background theme colour
function background() {
	var bg = localStorage.getItem("background");
	if (bg != undefined) {
		console.log("Setting background:", bg);
		document.body.style.backgroundColor = bg;
	}
	var click = function(id, colour) {
		console.log("Setting background:", id, colour);
		var item = document.getElementById(id);
		item.addEventListener("click", function() {
			document.body.style.backgroundColor = colour;
			localStorage.setItem("background", colour);
		});
	}
	click("grey", "silver");
	click("white", "whitesmoke");
	click("yellow", "beige");
}

function chars() {
	var i = document.getElementById('text').value.length;
	var c = maxChars;

	if (i > maxChars) {
		c = i - maxChars;
	} else {
		c = maxChars - i;
	}

	document.getElementById('chars').innerHTML = c;
};

function escapeHTML(str) {
	var div = document.createElement('div');
	div.style.display = 'none';
	div.appendChild(document.createTextNode(str));
	return div.innerHTML;
};

// call the backend api
function callAPI(ep, req, fn, errFn) {
	const http = new XMLHttpRequest();
	const url = "/api/" + ep;

	// when there is a change 
	http.onreadystatechange = function() {
		if (this.readyState == 4 && this.status == 200) {
			var data = JSON.parse(this.responseText);
			fn(data);
		} else if (this.status == 500) {
			console.log(ep, errFn);
			if (errFn != undefined) {
				errFn(this);
				return
			}
			var data = JSON.parse(this.responseText);

			if (data.error.charAt(0) == "{") {
				var err = JSON.parse(data.error)
				var error = document.getElementById("error");
				error.innerText = err.Detail;
			} else {
				var error = document.getElementById("error");
				error.innerText = data.error;
			}
		}
	}

	http.open("POST", url, true);
	http.setRequestHeader("Content-Type", "application/json;charset=UTF-8");
	http.send(JSON.stringify(req));
}

function delCookie(cname) {
	setCookie(cname, "", 0);
}

function getCookie(cname) {
	let name = cname + "=";
	let decodedCookie = decodeURIComponent(document.cookie);
	let ca = decodedCookie.split(';');

	for(let i = 0; i <ca.length; i++) {
		let c = ca[i];
		while (c.charAt(0) == ' ') {
		c = c.substring(1);
			}
		if (c.indexOf(name) == 0) {
			return c.substring(name.length, c.length);
		}
	}

	return "";
}

function setCookie(cname, cvalue, expiry) {
	const d = new Date(expiry * 1000);
	let expires = "expires="+ d.toUTCString();
	document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}

function timeSince(timestamp) {
	var ago = function(i, d) {
		var num = Math.floor(i);
		if (num <= 1) {
			return "1 " + d;
		}
		return num + " " + d + "s";
	}

	var date = Date.parse(timestamp);

	var seconds = Math.floor((new Date() - date) / 1000);

	var interval = seconds / 31536000;

	if (interval > 1) {
		return ago(interval, "year")
	}

	interval = seconds / 2592000;
	if (interval > 1) {
		return ago(interval, "month");
	}
	interval = seconds / 86400;
	if (interval > 1) {
		return ago(interval, "day");
	}

	interval = seconds / 3600;
	if (interval > 1) {
		return ago(interval, "hour");
	}

	interval = seconds / 60;
	if (interval > 1) {
		return ago(interval, "minute");
	}
	return ago(interval, "second");
}

function getVotes(post) {
	// add votes
	var upvotes = 0;
	var downvotes = 0;

	if (post.upvotes != undefined) {
		upvotes = post.upvotes;
	}
	
	if (post.downvotes != undefined) {
		downvotes = post.downvotes;
	}

	return upvotes - downvotes;
}

function getLinks(post) {
	// below title nav
	var links = document.createElement("span");
	links.style.fontSize = "small";
	links.setAttribute("class", "post-nav");

	if (post.url.length > 0) {
		links.innerHTML = "<a href='"+ post.url + "'>Link 🔗</a> | ";
	}
	var count = 0;

	// add comment count
	if (post.commentCount === undefined) {
		count = 0;
	} else {
		count = post.commentCount;
	}

	links.innerHTML += "<a href='#post="+ post.id + "'> 🗨️ (" + count + ")</a> ";

	// add votes
	var votes = getVotes(post);
	links.innerHTML += " | <a id='upvote-"+post.id+"' data-votes=" + votes + " href='#upvote=" 
			+ post.id + "' onclick='upvote(\""+post.id+"\"); return false;'><span class=upvote>🔺</span>" 
			+ votes + "</a>";

	return links;
}

// load the groups to list
function loadGroups() {
	var sidebar = document.getElementById("sidebar");
	var groups = document.createElement("div");
	groups.setAttribute("class", "groups");

	callAPI("groups", {}, function(data) {
		sidebar.innerHTML = '';
		groups.innerHTML = '';

		var newGroup = document.createElement("a")
		newGroup.setAttribute("class", "link right");
		newGroup.href = '#new-group';
		newGroup.innerText = 'New';
		sidebar.appendChild(newGroup);

		// set title
		var el = document.createElement("h4");
		el.innerHTML = '<a href="/#all">#️⃣</a>';
		el.setAttribute("class", "groups-title");
		sidebar.appendChild(el);

		if (data.records === undefined) {
			return
		}

		if (data.records.length == 0) {
			return
		}

		var selected;

		for (i = 0; i < data.records.length; i++) {
			var el = document.createElement("div");
			el.setAttribute("class", "group");
			var group = data.records[i];
			var hash = "#" + group.name.replaceAll(" ", "+");
			var button = document.createElement("a")
			button.href = hash;

			// set the title
			var title = document.createElement("h4");
			title.innerText = group.name.replace("m3o", "M3O");
			title.id = "group-" + group.id
			title.setAttribute("class", "group-title");

			var info = document.createElement("span");
			info.style.fontSize = "small";
			info.setAttribute("class", "group-info");
			info.innerText = group.description;

			var div = document.createElement("div");

			// set the selected attribute based on hash
			if (window.location.hash == hash) {
				el.setAttribute("class", "group selected");
				selected = el;
			}

			// highlight the current group selected
			el.onclick = function() {
				if (selected != undefined) {
					selected.setAttribute("class", "group");
				}
				this.setAttribute("class", "group selected");
				selected = this;
			}

			div.appendChild(title);
			div.appendChild(info);
			button.appendChild(div);

			el.appendChild(button);
			
			// append the group to groups list
			groups.appendChild(el);

			sidebar.appendChild(groups);
		}
	});
}

// load all the posts for a group
function loadGroup(name) {
	// don't process zero length names
	if (name.length == 0) {
		return;
	}

	// lowercase the name
	name = name.toLowerCase();

	var content = document.getElementById("content");
	// clear the content
	content.innerHTML = "";

	var newPost = document.createElement("div");
	newPost.setAttribute("id", "post-box");
	newPost.innerHTML += '<form id="new-post" action="#new-post" onsubmit="newPost(true); return false">'
	+ '<input id="title" name="title" placeholder="Topic" maxlength="256" required></input>'
	+ '<textarea id="text" name="text" placeholder="Message" maxlength="1024" required></textarea>'
	+ '<input id="group" name="group" type="hidden" value="'+ name +'"></input>'
	+ '<span id="chars">1024</span>'
	+ '<button class="link">Publish</button>'
	+ '</form>';
	content.appendChild(newPost);

	document.getElementById("text").addEventListener("keyup", function() {
		chars();
	});

	callAPI("posts", {"group": name }, function(data) {
		console.log("Got the data for " +  name + ": ",  data);
		if (data.records === undefined || data.records.length == 0) {
			content.innerHTML += "<p>Be the first to post something 👇</p>";
			return;
		}

		var posts = document.createElement("div");
		posts.setAttribute("class", "posts");

		for (i = 0; i < data.records.length; i++) {
			var el = document.createElement("div");
			el.setAttribute("class", "post");
			var post = data.records[i];

			// set the title
			var title = document.createElement("h4");
			title.innerHTML = "<a href='#post=" + post.id + "'><b>" + post.title + "</b></a>";
			title.id = "post-" + post.id
			title.setAttribute("class", "post-title");

			var info = document.createElement("span");
			info.style.fontSize = "small";
			info.setAttribute("class", "post-info");

			// posted by
			info.innerHTML = "By " + post.username + " " + timeSince(post.created) + " ago";

			var text = document.createElement("p");
			text.innerHTML = escapeHTML(post.content).parseURL(true).parseHashTag();
			text.id = "post-content-"+post.id;
			text.setAttribute("class", "post-content")
			
			// reducethe post length
			//if (post.content.length > 0) {
			//	text.innerHTML = escapeHTML(post.content.substring(0, 160));
			//	if (post.content.length > 160) {
			//		text.innerHTML += '...<a href="#post=' + post.id + '"><b><small>read more</small></b></a>';
			//	}
			//}

			// add group name if all
			if (name == "all") {
				var a = "<a class=group-link href='#" + post.group.replaceAll(" ", "+") + "'>" + post.group.replace("m3o", "M3O") + "</a>";
				info.innerHTML = info.innerHTML + " in " + a;
			}

			// create the nav links
			var links = getLinks(post);

			// append the content
			el.appendChild(title);
			el.appendChild(info);
			el.appendChild(text)
			el.appendChild(links);
			posts.appendChild(el);
		}

		content.appendChild(posts);
	});
}

// load the login page
function loadLogin() {
	var content = document.getElementById("content");

	content.innerHTML = `
		<h3 style="margin-top: 0">Login</h3>
		<form id="login-form", action="#login" onsubmit="login(true); return false">
		<p>
		  <input id="username" name="username" placeholder="Username" type=text minlength="1" required />
		</p>
		<p>
		  <input id="password" name="password" type="password" placeholder="Password" minlength="8" required />
		</p>
		<button class="link">Submit</button>
		</form>
		<h3>Signup</h3>
		<form id="signup-form", action="#signup" onsubmit="signup(true); return false">
		<p>
		  <input id="username" name="username" placeholder="Username" type=text minlength="1" required />
		</p>
		<p>
		  <input id="password" name="password" type="password" placeholder="Password" minlength="8" required />
		</p>
		<p>
		  <input id="email" name="email" type="email" placeholder="Email" required />
		</p>
		<button class="link">Submit</button>
		</form>
	`;
}

function loadPost(id) {
	callAPI("posts", { "id": id }, function(rsp) {
		var content = document.getElementById("content");
		// clear content
		content.innerHTML = "";

		var div = document.createElement("div");
		div.setAttribute("class", "post");

		if (rsp.records.length == 0) {
			content.innerText = "Post not found";
			return
		}

		var post = rsp.records[0];

		// set the title
		var title = document.createElement("h4");
		title.innerHTML = "<a href='#post=" + post.id + "'><b>" + post.title + "</b></a>";
		title.id = "post-" + post.id
		title.setAttribute("class", "post-title");

		// post content
		var p = document.createElement("p");
		p.setAttribute("class", "post-content")
		p.innerHTML = escapeHTML(post.content).parseURL(true).parseHashTag();

		// posted by
		var info = document.createElement("p");

		info.innerHTML += "By " + post.username + " " + timeSince(post.created) + " ago";
		info.style.fontSize = "small";

		var groupHash = "#" + post.group.replaceAll(" ", "+");

		// add group name
		var a = "<a class=group-link href='" + groupHash + "'>" + post.group.replace("m3o", "M3O") + "</a>";
		info.innerHTML = info.innerHTML + " in " + a;

		div.appendChild(title);
		div.appendChild(info);
		div.appendChild(p);
		content.appendChild(div);

		// comments
		var comments = document.createElement("div");
		comments.setAttribute("class", "comments");
		comments.innerHTML = 'Discussion';
		var link = document.createElement("a");
		var count = 0;

		if (post.commentCount === undefined) {
			count = 0;
		} else {
			count = post.commentCount;
		}

		if (count > 0) {
			callAPI("comments", {"postId": post.id }, function(rsp) {
				if (rsp.records.length == 0) {
					return
				}

				for (i = 0; i < rsp.records.length; i++) {
					var comment = rsp.records[i];

					var c = document.createElement("div");
					c.setAttribute("class", "comment");
					c.innerHTML = "<div><small>" + comment.username + " " + timeSince(comment.created) + " ago</small></div>";
					c.innerHTML += "<div>"+escapeHTML(comment.content).parseURL(true).parseHashTag()+"</div>";
					comments.appendChild(c);
				};

				console.log(rsp.records);
			});
		} else {
			comments.innerHTML = "";
		}

		var links = getLinks(post);
		div.appendChild(links);

		var comment = document.createElement("div");
		comment.id = 'comment-box';
		comment.innerHTML = `
			<form id="new-comment" action="#new-comment" onsubmit="newComment(true); return false;">
			  <textarea id="comment" cols=50 maxlength="500" rows=2 placeholder=Reply></textarea>
			  <input id="post" value="` + id + `" type=hidden />
			  <button class="link">Send</button>
			</form>
		`;
		content.appendChild(comment);
		content.appendChild(comments);

		//title.innerText = rsp.
		console.log(rsp);
	})
}

function loadSettings(update) {
	var session = getCookie("sess");

	if (session.length == 0) {
		console.log("bad session", session);
		return
	}

	if (update == true) {
		console.log("Updating settings");
		
		var el = document.getElementById("settings-form").elements;

		var sendEmails = el['emails'].checked;

		var settings = {
			"emails": sendEmails,
		};

		callAPI("settings", {"sessionId": session, "settings": settings}, function(rsp) {
			console.log("Settings saved");
		}, function(err) { console.log(err); })

		return;
	}


	callAPI("settings", {"sessionId": session}, function(rsp) {
		var emails = true;
		
		if (rsp != undefined) {
			emails = rsp["emails"];
		}
		console.log(emails);
		var content = document.getElementById("content");
		var form = document.createElement("form");
		form.id = "settings-form";
		form.action = "#settings";
		form.onsubmit = function(e) { e.preventDefault(); loadSettings(true); return false; };
		form.innerHTML = "Email Notifications:&nbsp;";

		var input = document.createElement("input");
		input.id = "emails";
		input.name = "emails";
		input.type = "checkbox";
		input.checked = emails;
		form.appendChild(input);
		form.appendChild(document.createElement("br"));

		var input = document.createElement("input");
		input.type = "submit";
		input.value = "update";
		input.setAttribute("class", "link")
		input.style.margin = "5px 0 0 0";
		form.appendChild(input)
		form.appendChild(document.createElement("br"));

		content.innerHTML = `<h3 style="margin-top: 0">Settings</h3>`;
		content.appendChild(form);
	}, function(err) { console.log(err) });
}

// display user info
function loadAccount(name) {
	var session = getCookie("sess");

	if (session.length == 0) {
		console.log("bad session", session);
		return
	}

	callAPI("readSession", {"sessionId": session}, function(rsp) {
		var content = document.getElementById("content");
		content.innerHTML = '<h3 style="margin-top:0">Account</h3>';

		var date = new Date(rsp.account.created * 1000);
		var joined = date.toLocaleDateString('en-us', {
			weekday:"long",
			month:"short",
			day:"numeric",
			year: "numeric",
		}); 


		var created = document.createElement("div");
		created.innerText = "Joined: " + joined;

		var verified = document.createElement("div");

		if (rsp.account.verified != undefined) {
			var verifiedDate = new Date(rsp.account.verification_date * 1000);
			verified.innerText = "Verified: " + verifiedDate.toLocaleDateString('en-us', {
				weekday:"long",
				month:"short",
				day:"numeric",
				year: "numeric",
			});
		} else {
			var verify = document.createElement("a");
			verify.setAttribute("class", "link");
			verify.id = "verify-email";
			verify.href = "#verify-email";
			verify.text = "Verify Account";
			verify.onclick = function(e) {
				e.preventDefault();
				callAPI("verifyAccount", {"sessionId": session}, function(rsp) {
					console.log(rsp)
				});
			};
			verified.appendChild(verify);
		}

		var username = document.createElement("div");
		username.innerText = "Username: " + rsp.account.username;

		var email = document.createElement("div");
		email.innerText = "Email: " + rsp.account.email;

		var balance = document.createElement("div");
		balance.innerText = "Balance: " + rsp.balance + " points";

		content.appendChild(username);
		content.appendChild(email);
		content.appendChild(balance);
		content.appendChild(created);
		content.appendChild(verified);
	}, function(err) { console.log(err) });
	
}

// login or authenticate the user
function login(submit) {
	// check if its a login
	if (submit == true) {
		console.log("Login event");
		
		var el = document.getElementById("login-form").elements;

		var username = el['username'].value;
		var password = el['password'].value;

		callAPI("login", {"username": username, "password": password }, function(rsp) {
			var expires = parseInt(rsp.session.expires)
			setCookie("sess", rsp.session.id, expires);

			window.location.href = "/";
			window.location.hash = "";
		}, function(err) { console.log(err); })

		return;
	}

	var session = getCookie("sess");

	if (session.length == 0) {
		console.log("bad session", session);
		return
	}

	callAPI("readSession", {"sessionId": session}, function(rsp) {
		var date = new Date();
		var now = Math.floor(Date.now() / 1000)
		var expires = parseInt(rsp.session.expires);

		// session expired
		if (expires < now) {
			return
		}

		var user = document.getElementById("user");
		user.innerText = "👤 " + rsp.account.username;
		// TODO: show balance == + " (" + rsp.balance + ")";

		var lg = document.getElementById("login-link");
		var st = document.getElementById("settings-link");
		user.style.display = 'inline-block';
		st.style.display = 'inline-block';
		st.innerText = "⚙️ Settings";
		lg.innerText = "🚪 Logout";
		lg.href = "#logout";
	}, function(err) { console.log(err) });
}

// logout the user
function logout() {
	var session = getCookie("sess");

	if (session.length == 0) {
		console.log("bad session", session);
		return;
	}

	callAPI("logout", {"sessionId": session}, function(rsp) {
		delCookie("sess");
		var user = document.getElementById("user");
		var lg = document.getElementById("login-link")
		var st = document.getElementById("settings-link");
		user.style.display = 'none';
		st.style.display = 'none';
		lg.innerText = "Login";
		lg.href = "#login";
		window.location.href = "/";
		window.location.hash = "";
	});
}

function newGroup(submit) {
	if (submit == true) {
		console.log("Group event");
		
		var session = getCookie("sess");

		var el = document.getElementById("new-group").elements;

		var name = el['name'].value;
		var description = el['description'].value;

		callAPI("group", {
			"group": {
				"name": name.toLowerCase(),
				"description": description,
			},
			"sessionId": session,
		}, function(rsp) {
			name = name.replaceAll(" ", "+");
			window.location.href = "/#" + name.toLowerCase();
		});

		return;
	}

	// render the form
	var content = document.getElementById("content");
	content.innerHTML = `
	        <h3 style="margin-top: 5px;">New #️⃣</h3>
		<form id="new-group" action="#new-group" onsubmit="newGroup(true); return false">
		  <p>
		    <input id="name" name="name" placeholder="Name" type=text minlength="1" required/>
		  </p>
		  <p>
		    <input id="description" name="description" placeholder="Description" type=text />
		  </p>
		  <button class="link">Submit</button>
		</form>
	`;
}

function newComment(submit) {
	if (submit == true) {
		console.log("Comment event");
		
		var session = getCookie("sess");

		var el = document.getElementById("new-comment").elements;

		var comment = el['comment'].value;
		var post = el["post"].value;

		callAPI("comment", {
			"comment": {
				"content": comment,
				"postId": post,
			},
			"sessionId": session,
		}, function(rsp) {
			loadPost(post);
		});

		return;
	}
}

function newPost(submit) {
	if (submit == true) {
		console.log("Post event");
		
		var session = getCookie("sess");

		var el = document.getElementById("new-post").elements;

		var title = el['title'].value;
		var group = el['group'].value;
		//var url = el['link'].value;
		var text = el['text'].value;

		callAPI("post", {
			"post": {
				"title": title,
				"group": group.toLowerCase(),
		//		"url": url,
				"content": text
			},
			"sessionId": session,
		}, function(rsp) {
			loadGroup(group);
		})

		return;
	}
}

// executes on hash reload and first load
function reload(refresh) {
	// clear the error on reload
	var error = document.getElementById("error");
	error.innerHTML = "";

	// clear content
	var content = document.getElementById("content");
	content.innerHTML = "";

	var hash = window.location.hash;

	// get the group name
	var name = hash.substring(1);
	var parts = name.split("=");
	var id = "";

	// e.g post=uuid
	if (parts.length == 2) {
		name = parts[0];
		id = parts[1];
	} else {
		name = name.replaceAll("+", " ");
	}

	var heading = name.replaceAll("-", " ");

	// get the route
	var route = routes.get(name);

	// load the groups if told to
	if (refresh == true) {
		// load groups
		loadGroups();
	}

	if (route != undefined) {
		setTitle("#" + name);
		console.log("Loading route: " + name);
		route(id);
		return
	}

	setTitle(hash);

	// check hash length
	if (hash.length == 0) {
		// set group to all
		name = "all";
		setTitle("#home");
	}


	// turn the name into the page title
	// load the route
	console.log("Loading group: " + name);

	// load the group
	loadGroup(name);
}

function setTitle(name) {
	document.title = name;
	var title = document.getElementById("title");
	title.innerText = name;
}

function signup(submit) {
	// check if its a login
	if (submit == true) {
		console.log("Signup event");
		
		var el = document.getElementById("signup-form").elements;

		var username = el['username'].value;
		var password = el['password'].value;
		var email = el['email'].value;

		callAPI("signup", {"username": username, "password": password, "email": email }, function(rsp) {
			setCookie("sess", rsp.session.id, rsp.session.expires);

			window.location.href = "/";
			window.location.hash = "";
		})

		return;
	}
}

function upvote(id) {
	var session = getCookie("sess");

	var up = document.getElementById("upvote-"+id);
	var votes = parseInt(up.getAttribute("data-votes"));

	votes += 1;
	up.setAttribute("data-votes", votes);
	up.innerHTML = '<span class=upvote>🔺</span>' + votes;

	callAPI("upvotePost", {"id": id, sessionId: session}, function(rsp) {
		console.log("upvoted", id);
	}, function(err) {
		console.log(err);
		votes -= 1;  
		up.setAttribute("data-votes", votes);
		up.innerHTML = '<span class=upvote>🔺</span>' + votes;
	})
}

function welcome(display) {
	var w = document.getElementById("welcome");

	if (display == false) {
		w.style.display = 'none';
		return;
	}

	var date = new Date();
	var hours = date.getHours();

	var time = date.toLocaleDateString('en-us', {
		weekday:"long",
		month:"short",
		day:"numeric",
		hour: "numeric",
		minute: "numeric",
	}); 

	var text = 'Welcome 🏠';

	if (hours < 12) {
		text = 'Good morning 🌞';
	} else if (hours < 18) {
		text = 'Good afternoon 🌥️';
	} else if (hours < 24) {
		text = 'Good evening 🌃';
	}

	w.innerHTML = text;
	w.innerHTML += '<span id="date">'+time+'</span>';
	w.style.display = 'block';
}

// the global router
var routes = new Map();
routes.set("login", loadLogin);
routes.set("logout", logout);
routes.set("new-group", newGroup);
routes.set("post", loadPost);
routes.set("settings", loadSettings);
routes.set("account", loadAccount);

// when the page is ready, start loading content
document.addEventListener("DOMContentLoaded", function(event) {
	login();
	reload(true);
	background();
})

window.addEventListener('hashchange', function() {
	reload(false);
});

