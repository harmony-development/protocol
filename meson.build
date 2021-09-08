project('harmony-protocol', 'c')

pkg = import('pkgconfig')

harmony = get_option('includedir') / 'harmony-protocols'

install_subdir('auth', install_dir: harmony)
install_subdir('batch', install_dir: harmony)
install_subdir('chat', install_dir: harmony)
install_subdir('emote', install_dir: harmony)
install_subdir('harmonytypes', install_dir: harmony)
install_subdir('mediaproxy', install_dir: harmony)
install_subdir('profile', install_dir: harmony)
install_subdir('sync', install_dir: harmony)
install_subdir('voice', install_dir: harmony)

lib = library('harmony-protocol')

pkg.generate(lib)
