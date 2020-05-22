
SCRIPTS := scripts

setup:
	python $(SCRIPTS)/gendata.py
	sudo mysql -u root < $(SCRIPTS)/setup.sql
	mysql courselect -u courselect -p < $(SCRIPTS)/schema.sql
	mysql courselect -u courselect -p < $(SCRIPTS)/testdata.sql