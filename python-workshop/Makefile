# Makefile

.PHONY: all venv run

all: venv run

venv:
	python3 -m venv .venv && \
	. .venv/bin/activate && \
	pip3 install -r requirements.txt

run:
	. .venv/bin/activate && \
	python3 app.py
