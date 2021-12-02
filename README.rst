.. _fhez: https://gitlab.com/deepcypher/python-fhez.git
.. |fhez| replace:: Python-FHEz

.. _go: https://go.dev/doc/
.. |go| replace:: go

.. _dc: https://deepcypher.me
.. |dc| replace:: DC

|dc|_ DarkLantern
=================

A lantern is a portable case that protects light, A dark lantern is one who's light can be hidden at will.
|dc|_ DarkLantern is a golang implementation of both deep learning and FHE to protect your precious secrets, while still guiding your way through difficult problems.

This library will be a sister project to |fhez|_, except for the |go|_ (golang) programming language.

FHE is a way in which we can process cyphertexts without ever decrypting them. Deep learning is a category of ways we can process data using neural network abstractions, often becoming state-of-the-art in any field that holds sufficient data with which to train the neural networks. Combining the two we propose, and strive for fully-open-source/  Kerckoffian, encrypted deep learning as a service.

Cypherpunks write code.

Testing
-------

To test this library during development we use adjacent `x_test.go` files. Since these will span multiple subdirectories we must give got test a sub-dir wildcard.

.. code-block:: bash
  :caption: Exhaustive sub-dir testing

  go test -v --cover ./...
