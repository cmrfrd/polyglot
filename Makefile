.NOTPARALLEL:

DOCKER? = docker
SHELL := /bin/bash
ALL_PROBLEMS = 1
PROBLEM_DIRECTORIES = $(wildcard ./problems/*/)
eq = $(and $(findstring $(1),$(2)),$(findstring $(2),$(1)))

## Check if paramater exists
ifdef PROBLEM

        $(info PROBLEM set to $(PROBLEM) )
        PROBLEM_DIRECTORIES = $(wildcard ./problems/$(PROBLEM)/*)

        ## Ensure problem directory exists
        ifneq "$(wildcard ./problems/$(PROBLEM)/ )" ""
                $(info Found problem directory $(PROBLEM) ...)
        else
                $(error Could not find problem directory $(PROBLEM) ...)
        endif
endif

ifdef LANGUAGE

        $(info LANGUAGE set to $(LANGUAGE) )

	ifndef PROBLEM

                PROBLEM_DIRECTORIES = $(wildcard ./problems/**/$(LANGUAGE))

                ## Ensure problem directory exists
                ifneq "$(wildcard ./problems/**/$(LANGUAGE) )" ""
                        $(info Found language problem directories $(PROBLEM_DIRECTORIES) ...)
                else
                        $(error Could not find language problem directories ...)
                endif
         else
                PROBLEM_DIRECTORIES = $(wildcard ./problems/$(PROBLEM)/$(LANGUAGE))

                ## Ensure problem directory exists
                ifneq "$(wildcard ./problems/$(PROBLEM)/$(LANGUAGE) )" ""
                        $(info Found language problem directory $(PROBLEM)/$(LANGUAGE) ...)
                else
                        $(error Could not find language problem directory $(PROBLEM)/$(LANGUAGE) ...)
                endif
         endif
endif

## Docker run function for individual language problem
define docker-run-problem-language
	@$(eval PROBLEM := $(1))
	@$(eval LANGUAGE := $(2))
	@$(eval ts := $(shell date +%s%N))
	@$(eval te := $(shell date +%s%N))

	@echo running problem $(PROBLEM) for language $(LANGUAGE);\
	ts=$$(date +%s%N) ;\
	docker run --rm --security-opt seccomp=unconfined -t $(PROBLEM)-$(LANGUAGE) ;\
	tt=$$((($$(date +%s%N) - $$ts)/1000000)) ;\
	echo "Time=$$tt (milli)"
endef

## Docker build for infividual language problem
define docker-build-problem-language
	@$(eval PROBLEM := $(1))
	@$(eval LANGUAGE := $(2))
	@$(eval CONTEXT := $(3))
	$(info building problem $(PROBLEM) for language $(LANGUAGE))
	$(if $(shell docker build -q -t $(PROBLEM)-$(LANGUAGE) $(3) > /dev/null ;\
		     [ $$? == 0 ] && echo || echo $$? ), \
	  $(error Unable to build problem $(PROBLEM) for language $(LANGUAGE)), \
	  $(info ))
endef

## Define language builder
define docker-mode-problem-language
	@$(eval MODE := $(1))
	@$(eval PROBLEM := $(shell basename $(shell dirname $(2))))
	@$(eval LANGUAGE := $(shell basename $(2)))

	$(if $(findstring $(MODE),build), $(call docker-build-problem-language, $(PROBLEM),$(LANGUAGE),$(2)))
	$(if $(findstring $(MODE),run), $(call docker-run-problem-language, $(PROBLEM),$(LANGUAGE)))
endef


##
## RULES
##

$(info Problem directories $(PROBLEM_DIRECTORIES))
PROBLEM_LANGUAGE_DIRECTORIES := $(foreach dir, ${PROBLEM_DIRECTORIES},  $(foreach lang, $(dir $(wildcard $(dir)*/*)), $(lang)) )

build: $(PROBLEM_LANGUAGE_DIRECTORIES)
	$(foreach PROBLEM_LANGUAGE, $^, $(call docker-mode-problem-language, build, $(shell pwd)/${PROBLEM_LANGUAGE}))

run: $(PROBLEM_LANGUAGE_DIRECTORIES)
	$(foreach PROBLEM_LANGUAGE, $^, $(call docker-mode-problem-language, run, $(shell pwd)/${PROBLEM_LANGUAGE}))

build-run: build run
