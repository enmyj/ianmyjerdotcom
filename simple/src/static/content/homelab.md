
---
### Creating a "Homelab"

I bought a Raspberry Pi at the beginning of Covid like everyone else but it's been
sitting in a box gathering dust. For some reason, I had a burst of inspiration (i.e., I
was watching reality TV with my partner and wanted something to do with my hands) and
decided to host my own website from my house. So, I flashed a new OS and tried
to set up something basic (but still reasonably robust) and different enough
technologies from work that I could learn some new techniques. Here's the rough setup
I'm using:

* **Kubernetes:** I use AWS ECS/Lambda and Docker a lot at work, so I wanted to try Kubernetes at home.
  I'm using `k3s` but I'm running my own `traefik` and I'm using `metalLB` instead of
  `serviceLB`
* **Cloudflare:** Cloudflare is my domain registrar which allows me to use [Cloudflare
  Tunnel](https://developers.cloudflare.com/cloudflare-one/connections/connect-networks/).
  I followed [this doc](https://developers.cloudflare.com/cloudflare-one/tutorials/many-cfd-one-tunnel/)
  to run the tunnel as a kubernetes workload. Cloudflare also automatically handles TLS
  certificates for `ianmyjer.com` and one-level of subdomains which is very nice!
* **Traefik:** I'm directing all Cloudflare Tunnel traffic to `traefik` because I wanted to learn how
  to use it. I'm using Traefik's `IngressRoute` for hostname and path prefix routing.
  I deployed `traefik` using Helm because the Helm Chart is maintained and was easy
  to follow.
* **DNS Sinkholes:** I got a bit carried away reading `r/selfhosted` and `r/homelab` and
  ended up deploying `pihole` as a Kubernetes workload for local ad
  blocking. Because I'm using `metalLB` instead of the built-in k3s ServiceLB, this
  kubernetes service has its own static LAN static IP within my home network. This
  allowed me to set it as a DNS server in my router's DHCP configuration, which has
  been particularly nice for blocking tracking requests from my "smart" TV. I opted to
  deploy `pihole` using a manually written Kubernetes Manifest because I didn't like the
  Helm chart.
* **Tailscale:** I got annoyed having to type `<local IP of Pihole>/admin/` into my
  browser so I could sigh in satisfaction that I had blocked 1% of the data collected
  about me. So, I started using the `tailscale` kubernetes operator with [cluster
  ingress routing](https://tailscale.com/kb/1439/kubernetes-operator-cluster-ingress).
  This means private services like Pihole and Traefik's admin dashboard are available
  (with HTTPS!) within my tailnet (but not publicly!). I'm also running tailscale
  directly on my raspberrypi itself.
* **Security:** Probably the most pleasing part of this setup is I didn't have to expose
  any router ports publicly. I can `ssh` to my raspberypi using my LAN static IP when
  I'm at home or via `tailscale ssh` when I'm not at home if necessary. It's a little
  crazy to me that Cloudflare and Tailscale are free for personal use like this. But,
  if the goal is for developers to bring these services from home to work, they have
  worked on me - their docs are great and the services have
  been nice to work with!

TL;DR - all "public" traffic (i.e., traffic to ianmyjer.com) goes Cloudflare -->
Cloudflare Tunnel --> Traefik --> Kubernetes Service. All "private" traffic is either by
LAN IP (courtesy of MetalLB) or via Tailscale. No router ports were opened in the making
of this "homelab" (i.e., thing I mess around with while watching reality TV lol)

The code and notes for this setup is here:
[https://github.com/enmyj/pi](https://github.com/enmyj/pi)

### Ian Myjer Dot Com

I use Javascript/React at work but it's not my strength and I didn't feel I needed
something so fancy to display basic information about myself. And, I use `Express.js` at
work and wanted to try out `go`. So, I decided to deploy this website using golang's
[Fiber](https://docs.gofiber.io/). I also saw a post on HackerNews about
[hellpot](https://github.com/yunginnanet/HellPot) which I thought was hilarious so a few
endpoints on `ianmyjer.com` are actually handled by hellpot instead 🍯.

The code is here:
[https://github.com/enmyj/ianmyjerdotcom](https://github.com/enmyj/ianmyjerdotcom)