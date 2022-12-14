#!/bin/bash

# Script for synchronizing bibliographies and other LaTeX files
# that are held in a directory
# "$HOME/Good-things/Bibliografia-i-pliki-LaTeXa"
# across the file system.

# !!!!!!!!!!
# For this reason you should always change ONE files in the directory
# # "$HOME/Good-things/Bibliografia-i-pliki-LaTeXa"


# Path to the directory with bibliographies and other LaTeX files
BIBLIOGRAPHYPATH="$HOME/GitHub-repositories/Writing-repositories/Błędy-i-uwagi/Nauki-rozne-bledy-i-uwagi/Bibliografia"
SPECIALLATEXFILESPATH="$HOME/GitHub-repositories/Writing-repositories/Błędy-i-uwagi/Nauki-rozne-bledy-i-uwagi/Pliki-LaTeXa"
TIKZFILESPATH="$HOME/GitHub-repositories/Writing-repositories/Błędy-i-uwagi/Nauki-rozne-bledy-i-uwagi/Pliki-PGFa-TikZa"
# Path to the main directory with LaTeX files containing errors found
# in books and other works and commentaries to these works.
ERRORSANDCOMMENTSPATH="$HOME/GitHub-repositories/Writing-repositories/Błędy-i-uwagi"
# Path to directory with list of excercises
EXERCISESLISTSPATH="$HOME/Good-things/Listy-zadań"


##############################
# Paths to the various directories in which bibliographies and other LaTeX
# files need be synchronized.

# ???? To trzeba uaktualnić, bo nie ma ścieżek do głównych repozytoriów
# Paths are divided in categories due to type of directories to which
# they lead. In one category you have directory with LaTeX files dealing
# with DEUS and philosophy in other with mathematics, etc.
# Every category starts with path of main directory, followed by paths
# to subdirectories in alphabetical order.
# With exception of category "DEUS and philosophy" all other categories
# are in the alphabetical order. Also path to subdirectories are listed
# in alphabetical order.



####################
# Path to main Git repositories of LaTeX files

# Path to the directory of functional analysis repository
FUNCTIONALANALYSISPATH="$ERRORSANDCOMMENTSPATH/Analiza-funkcjonalna-bledy-i-uwagi"

# Path to the directory of mathematical analysis repository
MATHEMATICALANALYSISPATH="$ERRORSANDCOMMENTSPATH/Analiza-matematyczna-bledy-i-uwagi"

# Path to the directory of various sciense repository
VARIOUSSCIENCESPATH="$ERRORSANDCOMMENTSPATH/Nauki-rozne-bledy-i-uwagi"

# Path to the directory of various works repository
VARIOUSWORKSPATH="$ERRORSANDCOMMENTSPATH/Rozne-dziela-bledy-i-uwagi"





####################
# DEUS and philosophy paths
DEUSPHILOSOPHYPATH="$VARIOUSSCIENCESPATH/DEUS-i-filozofia-błędy-i-uwagi/"
####################





####################
# Comics and graphic novels path

COMICSGRAPHICNOVELSPATH="$VARIOUSWORKSPATH/Komiksy-powieści-graficzne-błędy-i-uwagi"
####################





####################
# Computer science paths
COMPUTERSCIENCEPATH="$VARIOUSSCIENCESPATH/Informatyka-błędy-i-uwagi"

THEORYOFCOMPUTATIONPATH="$COMPUTERSCIENCEPATH/Teoria-obliczeń-błędy-i-uwagi"


##########
# Path to computer science exercises
COMPUTERSCIENCEEXERCISESPATH="$EXERCISESLISTSPATH/Informatyka-listy-zadań"
####################



####################
# Culture, analytics works path

CULTUREANALYSISPATH="$VARIOUSSCIENCESPATH/Pozostałe-dziedziny-błędy-i-uwagi/Kultura-błędy-i-uwagi"
####################



####################
# Economics paths
ECONOMICSPATH="$VARIOUSSCIENCESPATH/Pozostałe-dziedziny-błędy-i-uwagi/Ekonomia-błędy-i-uwagi"

# echo $ECONOMICSPATH
####################



####################
# Essays and journalism paths

ESSAYSANDJOURNALISMPATH="$VARIOUSWORKSPATH/Eseje-publicystyka-błędy-i-uwagi"
####################



#####################
# History paths
HISTORYPATH="$VARIOUSSCIENCESPATH/Historia-błędy-i-uwagi"
####################





#####################
# Lingustics paths
LINGUISTICSPATH="$VARIOUSSCIENCESPATH/Pozostałe-dziedziny-błędy-i-uwagi/Językoznawstwo-błędy-i-uwagi"
####################





####################
# Mathematics paths
MATHEMATICSPATH="$VARIOUSSCIENCESPATH/Matematyka-błędy-i-uwagi"


ALGEBRAPATH="$MATHEMATICSPATH/Algebra-błędy-i-uwagi"

ANALYSISANDNUMERICALMETHODSPATH="$MATHEMATICSPATH/Analiza-i-metody-numeryczne-błędy-i-uwagi"

ARITHMETICSANDNUMBERTHEORYPATH="$MATHEMATICSPATH/Arytmetyka-i-teoria-liczb-błędy-i-uwagi"

AUTOMATAANDFORMALLANGUAGESTHEORYPATH="$MATHEMATICSPATH/Teoria-automatów-i-języków-formalnych-błędy-i-uwagi"

CATEGORYTHEORYPATH="$MATHEMATICSPATH/Podstawy-matematyki-błędy-i-uwagi/Teoria-kategorii-błędy-i-uwagi"

COMBINATORICSPATH="$MATHEMATICSPATH/Analiza-kombinatoryczna-błędy-i-uwagi"

DIFFERENTIALEQUATIONSPATH="$MATHEMATICSPATH/Równania-różniczkowe-błędy-i-uwagi"

DIFFERENTIALGEOMETRYPATH="$MATHEMATICSPATH/Geometria-różniczkowa-błędy-i-uwagi"

LOGICANDSETTHEORYPATH="$MATHEMATICSPATH/Podstawy-matematyki-błędy-i-uwagi/Logika-i-teoria-mnogości-błędy-i-uwagi"

NONCOMMUTATIVEGEOMETRYPATH="$MATHEMATICSPATH/Geometria-nieprzemienna-błędy-i-uwagi"

PROBABILTYTHEORYPATH="$MATHEMATICSPATH/Rachunek-prawdopodobieństwa-błędy-i-uwagi"
##########

# Path to directory dedicated to books with mathematical exercises
MATHBOOKSWITHEXERCISESPATH="$MATHEMATICSPATH/Zbiory-zadań-z-matematyki-błędy-i-uwagi"

# Path to lists of mathematical exercises
MATHEXERCISESPATH="$EXERCISESLISTSPATH/Matematyka-listy-zadań"
####################





####################
# Exercises to do paths
EXERCISESTODOLISTPATH="$HOME/Good-things/Various-writings/Zadania-zrób"
####################





####################
# Physics paths
PHYSICSPATH="$VARIOUSSCIENCESPATH/Fizyka-błędy-i-uwagi"


ANALYSISOFEXPERIMENTALDATA="$PHYSICSPATH/Analiza-danych-eksperymentalnych-błędy-i-uwagi"

CLERKMAXWELLELECTRODYNAMICSPATH="$PHYSICSPATH/Elektrodynamika-Clerka-Maxwella-błędy-i-uwagi"

CONDENSEDMATTERPHYSICSPATH="$PHYSICSPATH/Fizyka-materii-skondensowanej-błędy-i-uwagi"

INTRODUCTIONSTOPHYSICSPATH="$PHYSICSPATH/Wprowadzenie-do-fizyki-błędy-i-uwagi"

MATHEMATICALPHYSICSPATH="$PHYSICSPATH/Fizyka-matematyczna-błędy-i-uwagi"

NEWTONIANMECHANICSPATH="$PHYSICSPATH/Mechanika-Newtona-błędy-i-uwagi"

PHYSICSANDOTHERDISCIPLINESPATH="$PHYSICSPATH/Fizyka-i-inne-dziedziny-błędy-i-uwagi"

PHYSICSBOOKSWITHEXERCISESPATH="$PHYSICSPATH/Zbiory-zadań-z-fizyki-błędy-i-uwagi"

PHYSICSOFSPACETIMEPATH="$PHYSICSPATH/Fizyka-czasoprzestrzeni-błędy-i-uwagi"

QFTPATH="$PHYSICSPATH/QFT-błędy-i-uwagi"

QUANTUMMECHANICSPATH="$PHYSICSPATH/Mechanika-kwantowa-błędy-i-uwagi"

TERMOSTATICSETCPATH="$PHYSICSPATH/Termostatyka-termodynamika-i-fizyka-statystyczna-błędy-i-uwagi"


##########
# Path to physics exercises

PHYSICSEXERCISESPATH="$EXERCISESLISTSPATH/Fizyka-listy-zadań"
####################




#####################
# Politology directory path
POLITOLOGYPATH="$VARIOUSSCIENCESPATH/Pozostałe-dziedziny-błędy-i-uwagi/Politologia-błędy-i-uwagi"
####################


####################
# Various books path

# VARIOUSBOOKSPATH="$ERRORSANDCOMMENTSPATH/Pozostałe-dziedziny-błędy-i-uwagi/Różne-książki-błędy-i-uwagi"
####################





####################
# Various temporary paths

NOTESONEPATH="$HOME/Good-things/Various-writings/Notatki-do-oddziaływań-elektrosłabych"

NOTESTWOPATH="$HOME/Good-things/Various-writings/Various-LaTeX"
####################





####################
# Various temporary paths

VERYSIMPLEBOOKSABOUTSCIENCEPATH="$VARIOUSSCIENCESPATH/Pozostałe-dziedziny-błędy-i-uwagi/Bardzo-proste-książki-o-nauce-błędy-i-uwagi"
####################





##############################
# Synchronization of bibliographies and various LaTeX files across
# the file system.
# Outside the cathegory "DEUS and philosophy", all other cathegories
# are in alphabetical order. Also subcategories inside cathegory



####################
# Synchronization of file "DEUSPhilBooks.bib"
LATEXFILEPATH="$BIBLIOGRAPHYPATH/DEUSPhilBooks.bib"


# DEUS and philosophy directory
rsync $LATEXFILEPATH $DEUSPHILOSOPHYPATH

# Mathematical analysis directory
rsync $LATEXFILEPATH $MATHEMATICALANALYSISPATH
####################



####################
# Synchronization of file "AudiovisualArtsBooks"
LATEXFILEPATH="$BIBLIOGRAPHYPATH/AudiovisualArtsBooks.bib"


# Works on culture directory
rsync $LATEXFILEPATH $CULTUREANALYSISPATH
####################



####################
# Synchronization of file "CivilizationCultureBooks.bib"
LATEXFILEPATH="$BIBLIOGRAPHYPATH/CivilizationCultureBooks.bib"


# Culture, analytics works directory
rsync $LATEXFILEPATH $CULTUREANALYSISPATH
####################



####################
# Synchronization of file "HistoryBooks.bib"
LATEXFILEPATH="$BIBLIOGRAPHYPATH/HistoryBooks.bib"


# History directory
rsync $LATEXFILEPATH $HISTORYPATH
####################



####################
# Synchronization of file "LiteraryStudiesBooks.bib"
LATEXFILEPATH="$BIBLIOGRAPHYPATH/LiteraryStudiesBooks.bib"


# Culture, analytics works directory
rsync $LATEXFILEPATH $CULTUREANALYSISPATH
####################



####################
# Synchronization of file "MathComScienceBooks.bib"
LATEXFILEPATH="$BIBLIOGRAPHYPATH/MathComScienceBooks.bib"



###############
# Computer science directories

# Computer science directory
rsync $LATEXFILEPATH $COMPUTERSCIENCEPATH

# Computer science excercises path
rsync $LATEXFILEPATH $COMPUTERSCIENCEEXERCISESPATH

# Theory of computation directory
rsync $LATEXFILEPATH $THEORYOFCOMPUTATIONPATH
###############



###############
# Mathematical directories

# Algebra directory
rsync $LATEXFILEPATH $ALGEBRAPATH

# Analysis and numerical methods directory
rsync $LATEXFILEPATH $ANALYSISANDNUMERICALMETHODSPATH

# Arithmetics and number theory directory
rsync $LATEXFILEPATH $ARITHMETICSANDNUMBERTHEORYPATH

# Automata and formal languages theory directory
rsync $LATEXFILEPATH $AUTOMATAANDFORMALLANGUAGESTHEORYPATH

# Category theory directory
rsync $LATEXFILEPATH $CATEGORYTHEORYPATH

# Combinatorics directory
rsync $LATEXFILEPATH $COMBINATORICSPATH

# Differential equations directory
rsync $LATEXFILEPATH $DIFFERENTIALEQUATIONSPATH

rsync $LATEXFILEPATH $DIFFERENTIALGEOMETRYPATH

rsync $LATEXFILEPATH $FUNCTIONALANALYSISPATH

# Logic and set theory directory
rsync $LATEXFILEPATH $LOGICANDSETTHEORYPATH

# Books with mathematical exercises directory
rsync $LATEXFILEPATH $MATHBOOKSWITHEXERCISESPATH

# Math excercises directory
rsync $LATEXFILEPATH $MATHEXERCISESPATH

rsync $LATEXFILEPATH $MATHEMATICALANALYSISPATH

rsync $LATEXFILEPATH $NONCOMMUTATIVEGEOMETRYPATH

rsync $LATEXFILEPATH $PROBABILTYTHEORYPATH
###############



# Very simple books about science directory
rsync $LATEXFILEPATH $VERYSIMPLEBOOKSABOUTSCIENCEPATH

####################



####################
# Synchronization of file "MusicBooks.bib"
LATEXFILEPATH="$BIBLIOGRAPHYPATH/MusicBooks.bib"


# Culture, analytics works directory
rsync $LATEXFILEPATH $CULTUREANALYSISPATH
####################



####################
# Synchronization of file "PhilNaturBooks.bib"
LATEXFILEPATH="$BIBLIOGRAPHYPATH/PhilNaturBooks.bib"



##########
# Mathematical directories

# Functional analysis directory
rsync $LATEXFILEPATH $FUNCTIONALANALYSISPATH
##########



##########
# Physics directories

# Introductions to physics directory
rsync $LATEXFILEPATH $INTRODUCTIONSTOPHYSICSPATH

# Clerk Maxwell electrodynamics
rsync $LATEXFILEPATH $CLERKMAXWELLELECTRODYNAMICSPATH

# Condensed matter physics path
rsync $LATEXFILEPATH $CONDENSEDMATTERPHYSICSPATH

# Mathematical physics directory
rsync $LATEXFILEPATH $MATHEMATICALPHYSICSPATH

# Newtonian mechanics directory
rsync $LATEXFILEPATH $NEWTONIANMECHANICSPATH

# Physics and other disciplines directory
rsync $LATEXFILEPATH $PHYSICSANDOTHERDISCIPLINESPATH

# Physics excercises directory
rsync $LATEXFILEPATH $PHYSICSEXERCISESPATH

# Books with physics excercise directory
rsync $LATEXFILEPATH $PHYSICSBOOKSWITHEXERCISESPATH

# Physics of spacetime directory
rsync $LATEXFILEPATH $PHYSICSOFSPACETIMEPATH

# QFT directory
rsync $LATEXFILEPATH $QFTPATH

# Quantum mechanics directory
rsync $LATEXFILEPATH $QUANTUMMECHANICSPATH

# Thermostatic, thermodynamics and statistical physics
rsync $LATEXFILEPATH $TERMOSTATICSETCPATH
##########



# Very simple books about science directory
rsync $LATEXFILEPATH $VERYSIMPLEBOOKSABOUTSCIENCEPATH
####################



####################
# Synchronization of file "VariousFieldsBooks.bib"
LATEXFILEPATH="$BIBLIOGRAPHYPATH/VariousFieldsBooks.bib"


# Culture, analytics works directory
rsync $LATEXFILEPATH $CULTUREANALYSISPATH

# Economics directory
rsync $LATEXFILEPATH $ECONOMICSPATH

# Economics directory
rsync $LATEXFILEPATH $ECONOMICSPATH

# Essays and jurnalism directory
rsync $LATEXFILEPATH $ESSAYSANDJOURNALISMPATH

# History directory
rsync $LATEXFILEPATH $HISTORYPATH

# Linguistic directory
rsync $LATEXFILEPATH $LINGUISTICSPATH

# Politology directory
rsync $LATEXFILEPATH $POLITOLOGYPATH

# Various books directory
rsync $LATEXFILEPATH $VARIOUSBOOKSPATH
####################



####################
# Synchronization of file "VisualArtsBooks.bib"
LATEXFILEPATH="$BIBLIOGRAPHYPATH/VisualArtsBooks.bib"


# # Economics directory
# rsync $LATEXFILEPATH $ECONOMICSPATH

# Comics and graphics novels directory
rsync $LATEXFILEPATH $COMICSGRAPHICNOVELSPATH

# Culture, analytics works directory
rsync $LATEXFILEPATH $CULTUREANALYSISPATH
####################













####################
# Synchronization of file "latexgeneralcommands.sty"
# It contains various general purpose macros.
LATEXFILEPATH="$SPECIALLATEXFILESPATH/latexgeneralcommands.sty"


# DEUS and philosophy directory
rsync $LATEXFILEPATH $DEUSPHILOSOPHYPATH



##########
# Computer science directories

# Computer science directory
rsync $LATEXFILEPATH $COMPUTERSCIENCEPATH

# Theory of computation directory
rsync $LATEXFILEPATH $THEORYOFCOMPUTATIONPATH
##########



# Culture, analytics works directory
rsync $LATEXFILEPATH $CULTUREANALYSISPATH

# Comics and graphics novels directory
rsync $LATEXFILEPATH $COMICSGRAPHICNOVELSPATH

# Computer science excercises direcotry
rsync $LATEXFILEPATH $COMPUTERSCIENCEEXERCISESPATH

# Economics directory
rsync $LATEXFILEPATH $ECONOMICSPATH

# Essays and jurnalism directory
rsync $LATEXFILEPATH $ESSAYSANDJOURNALISMPATH

# History directory
rsync $LATEXFILEPATH $HISTORYPATH

# Linguistic directory
rsync $LATEXFILEPATH $LINGUISTICSPATH

# Books with mathematical exercises directory
rsync $LATEXFILEPATH $MATHBOOKSWITHEXERCISESPATH

# Lists of math excercises directory
rsync $LATEXFILEPATH $MATHEXERCISESPATH

# Physics excercises directory
rsync $LATEXFILEPATH $PHYSICSEXERCISESPATH

# Politology directory
rsync $LATEXFILEPATH $POLITOLOGYPATH

# Various books directory
rsync $LATEXFILEPATH $VARIOUSBOOKSPATH

# Very simple books about science directory
rsync $LATEXFILEPATH $VERYSIMPLEBOOKSABOUTSCIENCEPATH



##########
# Mathematical directories

# Analysis and numerical methods directory
rsync $LATEXFILEPATH $ANALYSISANDNUMERICALMETHODSPATH

# Algebra path
rsync $LATEXFILEPATH $ALGEBRAPATH

# Arithmetics and number theory directory
rsync $LATEXFILEPATH $ARITHMETICSANDNUMBERTHEORYPATH

# Automata and formal languages theory directory
rsync $LATEXFILEPATH $AUTOMATAANDFORMALLANGUAGESTHEORYPATH

# Category theory directory
rsync $LATEXFILEPATH $CATEGORYTHEORYPATH

# Combinatorics directory
rsync $LATEXFILEPATH $COMBINATORICSPATH

# Differential equations directory
rsync $LATEXFILEPATH $DIFFERENTIALEQUATIONSPATH

# Differential geometry directory
rsync $LATEXFILEPATH $DIFFERENTIALGEOMETRYPATH

# Functional analysis directory
rsync $LATEXFILEPATH $FUNCTIONALANALYSISPATH

# Logic and set theory directory
rsync $LATEXFILEPATH $LOGICANDSETTHEORYPATH

# Mathematical analysis directory
rsync $LATEXFILEPATH $MATHEMATICALANALYSISPATH

# Noncommutative geometry directory
rsync $LATEXFILEPATH $NONCOMMUTATIVEGEOMETRYPATH

# Probability theory directory
rsync $LATEXFILEPATH $PROBABILTYTHEORYPATH
##########



##########
# Physics directories

# Analysis of experimental data directory
rsync $LATEXFILEPATH $ANALYSISOFEXPERIMENTALDATA

# Clerk Maxwell electrodynamics directory
rsync $LATEXFILEPATH $CLERKMAXWELLELECTRODYNAMICSPATH

# Condensed matter physics directory
rsync $LATEXFILEPATH $CONDENSEDMATTERPHYSICSPATH

# Introductions to physics directory
rsync $LATEXFILEPATH $INTRODUCTIONSTOPHYSICSPATH

# Mathematical physics directory
rsync $LATEXFILEPATH $MATHEMATICALPHYSICSPATH

# Newtonian mechanics directory
rsync $LATEXFILEPATH $NEWTONIANMECHANICSPATH

# Physics and other disciplines
rsync $LATEXFILEPATH $PHYSICSANDOTHERDISCIPLINESPATH

# Books with physics excercise directory
rsync $LATEXFILEPATH $PHYSICSBOOKSWITHEXERCISESPATH

# Physics of spacetime directory
rsync $LATEXFILEPATH $PHYSICSOFSPACETIMEPATH

# Quantum mechanics directory
rsync $LATEXFILEPATH $QUANTUMMECHANICSPATH

# QFT directory
rsync $LATEXFILEPATH $QFTPATH

# Thermostatic, thermodynamics and statistical physics
rsync $LATEXFILEPATH $TERMOSTATICSETCPATH
##########



# Various notes directories
rsync $LATEXFILEPATH $NOTESONEPATH

rsync $LATEXFILEPATH $NOTESTWOPATH



# Various books directory
rsync $LATEXFILEPATH $VARIOUSBOOKSPATH


####################



####################
# Synchronization of file "mathcommands.sty"
# It contains macros for mathematics formulas.
LATEXFILEPATH="$SPECIALLATEXFILESPATH/mathcommands.sty"


# DEUS and philosophy directory
rsync $LATEXFILEPATH $DEUSPHILOSOPHYPATH


##########
# Mathematical directories

# Algebra path
rsync $LATEXFILEPATH $ALGEBRAPATH

# Analysis and numerical methods directory
rsync $LATEXFILEPATH $ANALYSISANDNUMERICALMETHODSPATH

# Arithmetics and number theory directory
rsync $LATEXFILEPATH $ARITHMETICSANDNUMBERTHEORYPATH

# Automata and formal languages theory path
rsync $LATEXFILEPATH $AUTOMATAANDFORMALLANGUAGESTHEORYPATH

# Category theory directory
rsync $LATEXFILEPATH $CATEGORYTHEORYPATH

# Combinatorics directory
rsync $LATEXFILEPATH $COMBINATORICSPATH

# Books with mathematical exercises directory
rsync $LATEXFILEPATH $MATHBOOKSWITHEXERCISESPATH

# Differential equations directory
rsync $LATEXFILEPATH $DIFFERENTIALEQUATIONSPATH

# Differential geometry directory
rsync $LATEXFILEPATH $DIFFERENTIALGEOMETRYPATH

# Functional analysis path
rsync $LATEXFILEPATH $FUNCTIONALANALYSISPATH

# Logic and set theory directory
rsync $LATEXFILEPATH $LOGICANDSETTHEORYPATH

# Mathematical analysis directory
rsync $LATEXFILEPATH $MATHEMATICALANALYSISPATH

# Noncommutative geometry directory
rsync $LATEXFILEPATH $NONCOMMUTATIVEGEOMETRYPATH

# Probability theory directory
rsync $LATEXFILEPATH $PROBABILTYTHEORYPATH


##########
# Physics directories

# Analysis of experimental data directory
rsync $LATEXFILEPATH $ANALYSISOFEXPERIMENTALDATA

# Clerk Maxwell electrodynamics
rsync $LATEXFILEPATH $CLERKMAXWELLELECTRODYNAMICSPATH

# Condensed matter physics directory
rsync $LATEXFILEPATH $CONDENSEDMATTERPHYSICSPATH

# Introductions to physics directory
rsync $LATEXFILEPATH $INTRODUCTIONSTOPHYSICSPATH

# Mathematical physics directory
rsync $LATEXFILEPATH $MATHEMATICALPHYSICSPATH

# Newtonian mechanics directory
rsync $LATEXFILEPATH $NEWTONIANMECHANICSPATH

# Physics and other disciplines
rsync $LATEXFILEPATH $PHYSICSANDOTHERDISCIPLINESPATH

# Books with physics excercise directory
rsync $LATEXFILEPATH $PHYSICSBOOKSWITHEXERCISESPATH

# Physics of spacetime directory
rsync $LATEXFILEPATH $PHYSICSOFSPACETIMEPATH

# Quantum mechanics directory
rsync $LATEXFILEPATH $QUANTUMMECHANICSPATH

# QFT directory
rsync $LATEXFILEPATH $QFTPATH

# Thermostatic, thermodynamics and statistical physics
rsync $LATEXFILEPATH $TERMOSTATICSETCPATH


##########
# Various notes directories

rsync $LATEXFILEPATH $NOTESONEPATH
####################




####################
# Synchronization of file "functionalanalysiscommands.sty"
# It contains various macros for symbols that often appeare in the context
# of functional analysis
LATEXFILEPATH="$SPECIALLATEXFILESPATH/functionalanalysiscommands.sty"


rsync $LATEXFILEPATH $FUNCTIONALANALYSISPATH
####################





####################
# Synchronization of file "TikZPics.sty"
# It contains ????
LATEXFILEPATH="$TIKZFILESPATH/TikZPics.sty"


rsync $LATEXFILEPATH $FUNCTIONALANALYSISPATH
####################





####################
# Synchronization of file "TikZStyles.sty"
# It contains ????
LATEXFILEPATH="$TIKZFILESPATH/TikZStyles.sty"


rsync $LATEXFILEPATH $FUNCTIONALANALYSISPATH
####################


# ##########
