
---
Last updated: 2025-01-04

### Setting up Kubernetes and deploying a few things

I bought a Raspberry Pi at the beginning of Covid like everyone else, but it has been
gathering dust. Recently, I had a burst of inspiration (i.e., I
was watching reality TV with my partner and wanted something to do with my hands) and
decided I wanted to host `ianmyjer.com` from my house. So, I tried to set up something
basic (but still reasonably robust) using technologies I don't work with regularly so
I could learn something new. Here's the setup as it stands today:

* **Kubernetes:** I use AWS ECS/Lambda and Docker a lot at work, so I wanted to try
  Kubernetes at home. I'm using `k3s` but I'm running my own `traefik` and I'm using
  `metalLB` instead of `serviceLB`
* **Cloudflare:** Cloudflare is my domain registrar which allows me to use [Cloudflare
  Tunnel](https://developers.cloudflare.com/cloudflare-one/connections/connect-networks/).
  I followed [this doc](https://developers.cloudflare.com/cloudflare-one/tutorials/many-cfd-one-tunnel/)
  to run the tunnel as a kubernetes workload. Cloudflare automatically (and freely)
  handles TLS certificates for `ianmyjer.com` and some static asset caching which is
  pretty cool!
* **Traefik:** I'm directing all Cloudflare Tunnel traffic to `traefik` because I wanted
  to learn how to use it. I'm using Traefik's `IngressRoute` for hostname and path
  prefix routing. I deployed `traefik` using Helm because the Helm Chart is maintained
  and was easy to follow.
* **DNS Sinkholes:** I got a bit carried away reading `r/selfhosted` and `r/homelab` and
  ended up deploying `pihole` as a Kubernetes workload for local ad
  blocking (I also tried `blocky` and `adguard home` but pihole has the best
  documentation). Because I'm using `metalLB` instead of the built-in k3s `serviceLB`,
  `pihole` has its own LAN static IP within my home network. This
  allowed me to set it as a DNS server on my router's DHCP configuration, which has
  been particularly nice for blocking requests from my "smart" TV. I opted to
  deploy `pihole` using a manually written Kubernetes Manifest because I didn't like the
  Helm chart.
* **Tailscale:** I got annoyed having to type `<local IP of Pihole>/admin/` into my
  browser so I could sigh with satisfaction that I had blocked <1% of the data collected
  about me. So, I started using the `tailscale` kubernetes operator with [cluster
  ingress routing](https://tailscale.com/kb/1439/kubernetes-operator-cluster-ingress).
  This means private services like pihole and Traefik's admin dashboard are available
  (with HTTPS!) within my tailnet (but not publicly!). I'm also running tailscale
  directly on my raspberrypi itself.
* **Security:** Maybe the nicest part of this setup is Cloudflare Tunnel +
  Tailscale allowed me to not open any router ports publicly. I can `ssh` to my
  raspberypi using my LAN static IP when I'm at home or via `tailscale ssh` when I'm not
  at home if necessary. It's very nice that Cloudflare and Tailscale are free
  for personal use like this!
* **Professional Stuff:** I briefly had prometheus/grafana running via the Helm chart
  but they were using valuable resources on my raspberry pi and I'm never
  going to check the monitoring anyway so I took it down. I also thought about setting
  up CI/CD but I spend a lot of time with Github Actions at work so I backburnered that
  idea 🙃.

TL;DR - all "public" traffic goes Cloudflare --> Cloudflare Tunnel --> Traefik -->
Kubernetes Service. All "private" traffic is either to static LAN IPs (courtesy of
MetalLB) or via Tailscale. No router ports were opened in the making of this "homelab"

Code and notes for my setup:
[https://github.com/enmyj/pi](https://github.com/enmyj/pi)

### Ian Myjer Dot Com

I use Javascript/React at work but it's not my strength and I didn't feel I needed
something so fancy to convey basic information about myself (one could certainly look at
the sorry design of this site and argue exactly the opposite, but I did get good marks
on the Google PageSpeed Insights analyzer so 🤷). And, I use `Express.js` at work but
wanted to try out `go` so I decided to deploy this website using 
[Fiber](https://docs.gofiber.io/). Plus, their benchmarks say it's super fast which is
ultimately what's important for websites, right?!

I also saw a post on HackerNews about
[hellpot](https://github.com/yunginnanet/HellPot) which I thought was hilarious so a few
endpoints [https://ianmyjer.com/robots.txt](https://ianmyjer.com/robots.txt) are
actually handled by hellpot instead 🍯.

The code is here:
[https://github.com/enmyj/ianmyjerdotcom](https://github.com/enmyj/ianmyjerdotcom)