# Makefile

SHELL := /bin/bash
.DEFAULT_GOAL := all

.PHONY: all
all: venv run

.PHONY: venv
venv:
	python -m venv .venv && \
	source .venv/bin/activate &&\
	pip install -r requirements.txt

.PHONY: run
run:
	source .venv/bin/activate
	python app.py
