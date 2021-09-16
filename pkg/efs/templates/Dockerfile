FROM public.ecr.aws/lambda/provided:al2
ADD ./bin /var/task/
ADD ./stash /var/task/stash
ADD ./cstore.yml /var/task/cstore.yml
ENTRYPOINT ["/var/task/bin"]
