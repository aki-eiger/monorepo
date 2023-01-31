
APPLICATIONS=app1 app2

.PHONY=buildall runlocalall deployall

define ALL_APPS
	for APP in $(APPLICATIONS); do \
		$(MAKE) -C $$APP $(1); \
	done
endef

buildall:
	$(call ALL_APPS, build)

deployall:
	$(call ALL_APPS, deploy)

runlocalall:
	$(call ALL_APPS, runlocal)

	
