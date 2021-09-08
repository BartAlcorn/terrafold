FROM public.ecr.aws/lambda/provided:al2
ADD ./handler/invoke/APP /var/task/
ENTRYPOINT ["/var/task/APP"]