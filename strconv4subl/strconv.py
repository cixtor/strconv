import base64
import hashlib
import sublime
import sublime_plugin
import subprocess
import urllib.parse


class StrconvCommand(sublime_plugin.TextCommand):
	def run(self, edit, **args):
		selections = self.view.sel()
		if len(selections) == 0:
			return
		for region in selections:
			the_text = self.view.substr(region)
			if the_text == "":
				self.set_status("nothing to format")
				continue
			if args.get("action") == "md5":
				the_text = hashlib.md5(the_text.encode("utf-8")).hexdigest()
			elif args.get("action") == "sha1":
				the_text = hashlib.sha1(the_text.encode("utf-8")).hexdigest()
			elif args.get("action") == "b64enc":
				the_text = base64.b64encode(the_text.encode("utf-8")).decode("utf-8")
			elif args.get("action") == "b64dec":
				the_text = base64.b64decode(the_text.encode("utf-8")).decode("utf-8")
			elif args.get("action") == "urlenc":
				the_text = urllib.parse.quote(the_text)
			elif args.get("action") == "urldec":
				the_text = urllib.parse.unquote(the_text)
			elif args.get("action") == "rgbfy":
				the_text = self.rgbfy(the_text)
			else:
				the_text = self.external_command(the_text, **args)
			region_text = self.view.substr(region)
			self.view.replace(edit, region, the_text)
			self.set_status("done")


	def external_command(self, the_text, **args):
		p = subprocess.Popen(
			["strconv", args.get("action")],
			stdin=subprocess.PIPE,
			stderr=subprocess.PIPE,
			stdout=subprocess.PIPE,
		)
		out, err = p.communicate(input=the_text)
		if p.returncode != 0:
			self.set_status(err.split("\n")[0])
			return the_text
		return out


	def rgbfy(self, the_text):
		h = hashlib.md5(the_text.encode("utf-8")).digest()
		return "#{:02X}{:02X}{:02X}".format(h[0], h[1], h[2])


	def set_status(self, message):
		self.view.set_status("strconv_output", message)
		sublime.set_timeout_async(
			lambda: self.view.erase_status("strconv_output"), 5000
		)
