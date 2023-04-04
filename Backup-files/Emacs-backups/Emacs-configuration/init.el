;; #################
;; Load the `config.org' file
(org-babel-load-file (expand-file-name "config.org" user-emacs-directory))


;; #####
;; Load `configadd.org' file
;; Additiona configuration for Emacs
(org-babel-load-file (expand-file-name "configadditional.org"
				       user-emacs-directory))

;; #####
;; Load `configovpowered.org' file
;; With configuraiotn of problematic and "too useful" Emacs
;; setting and packages.
(org-babel-load-file (expand-file-name "configovpowered.org"
				       user-emacs-directory))







;; #################
;; Cofiguration of ido (Interactive DO things)

;; ;; Maybe this is still needed. But I don't know.
;; (setq ido-file-extensions-order '(".tex" ".org" ".txt" ".el"
;; 				  ".rs" ".lisp"))
;; (setq ido-ignore-extensions t) 		; Ignore objects definde by
;; ;; `complete-ignored-extensions' variable





;; #################
;; Configuration of buffers list
;; (global-set-key (kbd "C-x C-b") 'ibuffer-other-window)

;; (setq ibuffer-default-sorting-mode 'major-mode)  ; Sorts buffers
;; ;; in buffers list by major mode of buffer
