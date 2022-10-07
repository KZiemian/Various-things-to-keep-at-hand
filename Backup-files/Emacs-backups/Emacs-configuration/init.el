;; #################
;; First steps of initialization
;; Set up `windows-system' at the begining to avoid short displaying of
;; it default look at initialization (Daniel Mai say so)


;; #########
;; Configuration of frames
(when window-system
  (menu-bar-mode -1)
  ;; Evaluate with positive integer to show menu in the top of frame.
  ;; (menu-bar-mode 1)
  (tool-bar-mode -1)
  (scroll-bar-mode -1)
  (tooltip-mode -1)  ; I don't know what this doing, but I never use
  ;; tooltips, so this probably doesn't matter.
  (set-frame-size (selected-frame) 80 100))


(setq inhibit-startup-message t)  ; Disable start up message.





;; #################
;; Configuration of package menager
(require 'package)
(setq package-enable-at-startup nil)

;; Setting up list of repositories
(add-to-list 'package-archives
	     '("melpa" . "https://melpa.org/packages/") t)
(add-to-list 'package-archives
	     '("gnu" . "https://elpa.gnu.org/packages/") t)
(add-to-list 'package-archives
	     '("marmelade" . "https://marmalade-repo.org/packages/") t)
(add-to-list 'package-archives
	     '("org" . "http://orgmode.org/elpa/") t)  ; This may be
;; redundant, but for "spokoju sumiania" I left if here

;; Disable package signature
(setq package-check-signature nil)

;; ??????
(when (boundp 'package-pinned-packages)
  (setq package-pinned-packages
	'((org-plus-contrib . "org"))))


;; ???????
(package-initialize)





;; #################
;; `Use-package' -- bookstraping package to manage packages
(unless (package-installed-p 'use-package)
  (package-refresh-contents)
  (package-install 'use-package))

(setq use-package-verbose t
      load-prefer-newer t)


;; From `use-package' README
(eval-when-compile
  (require 'use-package))
(require 'bind-key)





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
