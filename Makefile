
SCRIPTS := scripts

setup:
	sudo mysql -u root < $(SCRIPTS)/setup.sql
	sudo mysql courselect -u courselect -p < $(SCRIPTS)/schema.sql