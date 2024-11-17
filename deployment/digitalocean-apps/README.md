# Capsa deployment DigitalOcean Apps

In this directory you can find what is perhaps the easiest way to set up Capsa: with DigitalOcean Apps.

Deployment is quite straight-forward: take the specs, set the correct values, add them in DigitalOcean and you should be up and running in a few minutes!

DigitalOcean will automatically add TLS certificates on the App Platform, once the app is connected to the domain name.

## Domains

If you don't have the domain for Capsa managed in DigitalOcean, remove the `domains` block, apply the specs, and then later manually set the DNS to link them together.
