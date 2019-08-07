
if [ -z "$GOMEETING_HOST" ]
then
  echo "You must export the \$GOMEETING_HOST for telling me to use which server !"
  echo "Example:"
  echo "export GOMEETING_HOST=localhost:7728"
  exit 1
fi

if [ -z "$GOMEETING_TOKEN" ]
then
  echo "You must export the \$GOMEETING_TOKEN for authentication !"
  echo "Example:"
  echo "export GOMEETING_TOKEN=xxxx"
  exit 1
fi