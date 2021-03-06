pack_file = GoMeeting-$(shell git tag -l | tail -n 1)-$(shell uname -s)-$(shell uname -p).tar.xz
css_dir = assets/css
js_dir = assets/js
tmp_dir = TEST/tmp
vendor_dir = vendor
bulma = $(css_dir)/bulma.min.css
bulma_src = $(tmp_dir)/bulma-0.7.5/css/bulma.min.css
laydate = $(js_dir)/laydate
laydate_src = $(tmp_dir)/laydate-5.0.9/dist
axios = $(js_dir)/axios.min.js
axios_src = $(tmp_dir)/axios-0.18.1/dist/axios.min.js
axios_map = $(js_dir)/axios.min.map
axios_map_src = $(tmp_dir)/axios-0.18.1/dist/axios.min.map
fontawesome = $(js_dir)/fontawesome.min.js
fontawesome_src = $(tmp_dir)/fontawesome-free-5.11.2-web/js/all.min.js
md5 = $(js_dir)/md5.min.js
md5_src = $(tmp_dir)/js-md5-0.7.3/build/md5.min.js
vue = $(js_dir)/vue.min.js
vue_src = $(tmp_dir)/vue-2.6.10/dist/vue.min.js
md2html = $(js_dir)/markdown.min.js
md2html_src = $(tmp_dir)/markdown-browser-0.6.0-beta1/markdown.min.js


define getSrc
	curl -fsSL -o $(tmp_dir)/tmp.zip  $(1)
	unzip -q $(tmp_dir)/tmp.zip -d $(tmp_dir)
	rm $(tmp_dir)/tmp.zip
endef

all: clean $(vendor_dir) $(tmp_dir) $(bulma) $(laydate) $(axios) $(axios_map) $(fontawesome) $(md5) $(vue) $(md2html)
	-mkdir bin
	go build -o bin/gomeeting -ldflags "-w" main.go
pack: all
	-mkdir TEST
	-@rm TEST/$(pack_file)
	tar cJf TEST/$(pack_file) assets bin script config.yml.sample
clean:
	-rm -rf bin
clean_all: clean
	-rm -rf $(tmp_dir) \
		$(vendor_dir) \
		$(bulma) \
		$(laydate) \
		$(axios) \
		$(axios_map) \
		$(fontawesome) \
		$(md5) \
		$(vue) \
		$(md2html)

$(vendor_dir):
	glide install
$(tmp_dir):
	-mkdir -p $(tmp_dir)

$(bulma): $(bulma_src)
	@cp -av $< $@
$(bulma_src):
	$(call getSrc,"https://github.com/jgthms/bulma/releases/download/0.7.5/bulma-0.7.5.zip")

$(laydate): $(laydate_src)
	@cp -av $< $@
$(laydate_src):
	$(call getSrc,"https://github.com/sentsin/laydate/archive/v5.0.9.zip")

$(axios): $(axios_src)
	@cp -av $< $@
$(axios_src):
	$(call getSrc,"https://github.com/axios/axios/archive/v0.18.1.zip")
$(axios_map): $(axios_map_src)
	@cp -av $< $@
$(axios_map_src):
	$(call getSrc,"https://github.com/axios/axios/archive/v0.18.1.zip")

$(fontawesome): $(fontawesome_src)
	@cp -av $< $@
$(fontawesome_src):
	$(call getSrc,"https://github.com/FortAwesome/Font-Awesome/releases/download/5.11.2/fontawesome-free-5.11.2-web.zip")

$(md5): $(md5_src)
	@cp -av $< $@
$(md5_src):
	$(call getSrc,"https://github.com/emn178/js-md5/archive/v0.7.3.zip")

$(vue): $(vue_src)
	@cp -av $< $@
$(vue_src):
	$(call getSrc,"https://github.com/vuejs/vue/archive/v2.6.10.zip")

$(md2html): $(md2html_src)
	@cp -av $< $@
$(md2html_src):
	curl -fsSL -o $(tmp_dir)/tmp.tgz 'https://github.com/evilstreak/markdown-js/releases/download/v0.6.0-beta1/markdown-browser-0.6.0-beta1.tgz'
	tar -xf $(tmp_dir)/tmp.tgz -C $(tmp_dir)
	rm $(tmp_dir)/tmp.tgz