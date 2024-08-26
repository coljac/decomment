" simple_function.vim

function! SimpleFunction()
echo "Hello from SimpleFunction"
endfunction

command! CallSimpleFunction call SimpleFunction()
nnoremap <leader>sf :CallSimpleFunction<CR>

" Usage:
" :CallSimpleFunction
" or
" Press <leader>sf in normal mode

