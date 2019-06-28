FROM fedora
RUN yum update -y &&    yum install -y sbcl
ADD solution.clisp /solution.clisp
CMD ["bash", "-c", "sbcl --load /solution.clisp | tail -n +8"]
