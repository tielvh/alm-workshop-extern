# ALM Workshop 2025

## Introduction

Hi everyone! Welcome to the **ALM Workshop 2025**!

This workshop is designed to give you hands-on experience with:
* Learning new and unfamiliar technologies.
* Implementing **Continuous Integration (CI) and Application Lifecycle Management (ALM)** concepts.
* Applying the theory you’ve learned in a real-world setup.

To make things more interesting, we’ll be using technologies that you may not have worked with before:
* **Golang** – A modern programming language known for its efficiency and simplicity.
* **Quay** – A container registry for storing and managing container images.
* **GitHub DevSpaces / DevContainers** – A development environment that works in the cloud.
* **GitHub Actions** – A CI/CD pipeline tool for automating software workflows.
* **RedHat OpenShift** – A Kubernetes-based container platform for deploying applications.

## Workshop Structure

The workshop is divided into several tasks, each serving as a checkpoint. If you get stuck or run out of time, you can always check out the corresponding branch to catch up.

---

## Prerequisites

Before starting, follow these steps:
1. **Fork** this repository to your own GitHub account.
2. Start a **GitHub CodeSpace** from the `main` branch.

    ![Create Codespace](docs/github-create-codespace.png)
    * This may take **2+ minutes**, so be patient.
    * You may get pop-ups asking to install extensions—accept them.

3. Review the repository structure:
    * **.devcontainer/devcontainer.json** – Configures our development environment.
    * **.github/workflows/** - Contains our yaml files for our CI/CD pipelines.
    * **infra/openshift.yaml** – Contains our Kubernetes resources.
    * **workshop-service/** – Contains the source code.
    * **Dockerfile** – A multi-stage build file for our Go application.


4. Test your Docker setup by building and running the image:
    ```sh
    docker build -t test-workshop . --progress=plain
    docker run -p 3000:3000 test-workshop
    ```
    * Your app should be accessible via the **GitHub workspace URL** or **localhost**.
      * If you missed the popup you can click on the little transmission tower in the bottom left

      ![VSCode Transmission Tower](docs/vscode-portforward.png)

    * Visit the `/workshop` endpoint to check if it works.
      * You can also share your links that GitHub codespaces generate when you're running your application, but by default the connection is set to private! 
      * You can change the visibilty in the same ports tab as before as you can see in the image.

      ![Change port visibility](docs/vscode-port-visiblity.png)

Now that we have a working development environment, let's move on!

---

## Tasks

### **Task 1: Setting Up a CI Pipeline**
#### **Problem Statement**
Currently, we only have a local development setup. Our goal is to:
* Work with multiple developers while ensuring **consistent, repeatable builds**.
* Follow **Software Development Lifecycle (SDLC)** best practices.
* Deploy our application publicly in a reliable manner.

#### **Solution Approach**
We'll set up a **GitHub Actions pipeline** to automate the build and deployment process.

#### **Steps**
1. A pre-made **GitHub Actions workflow** file for this part is already present inside `.github/workflows/ci.yaml`.
2. Modify the steps to **build** and **push** the container image to [**Quay.io**](https://quay.io/).
3. Use a [**Robot Account**](https://docs.quay.io/glossary/robot-accounts.html) to authenticate with Quay.io.
   * **HINT**: The image name should be as such quay.io/<your-username>/alm-workshop to properly push
   * **HINT**: You might need to register the repository in quay first before you can push.
   * **HINT**: How will you manage your quay Robot User password in the pipeline?
4. Validate the deployment by pulling the container image and running it locally:
    ```sh
    docker run -p 3000:3000 quay.io/<your-username>/alm-workshop:<tag>
    ```

---

### **Task 2: Implementing Semantic Versioning (SemVer)**
#### **Problem Statement**
Currently, all builds result in an unversioned image, making it hard to track releases.

#### **Solution Approach**
We'll implement **Semantic Versioning (SemVer)** to tag images correctly upon release.

This means that when you create a GitHub release with a certain SemVer, the pipeline should be ran again and the docker-image should be `<quay-user>/alm-workshop:x.x.x`

#### **Steps**
1. Modify the CI pipeline to trigger on **GitHub Releases**.
   * **HINT**: https://docs.github.com/en/actions/using-workflows/events-that-trigger-workflows#release
2. Create a **GitHub release** for version `1.0.0`.

---

### **Task 3: Deploying to OpenShift (Manual)**
#### **Problem Statement**
Right now, our application is only running locally. We need a **publicly accessible deployment**.

#### **Solution Approach**
We'll deploy our container to **RedHat OpenShift Sandbox**.

#### **Steps**
1. Log in to [OpenShift Sandbox](https://console.redhat.com/openshift/sandbox).
2. Press "Getting started"
3. You should be able to launch OpenShift
   
   ![OpenShift Launch](docs/redhat-start-page.png)
   
4. Go to the **Developer menu → +Add → Container Images** option.
5. Deploy your Quay.io container image by using the image registry URL.
   * Runtime icon = Golang
   * Deploy resource type = deployment
   * Rest should be OK by default
6. Verify deployment via **Topology View** and try to access the remote endpoint.
   * Click the little arrow icon above your application to open the route.
   
   ![OpenShift Route](docs/openshift-topology-open-route.png)

---

### **Task 4: Automating OpenShift Deployments**
#### **Problem Statement**
Manually deploying an application for every change is time-consuming and error-prone.

#### **Solution Approach**
We'll use **Infrastructure as Code (IaC)** to automate OpenShift deployments.

#### **Steps**
1. Delete the manual deployment from **Topology View**.

    ![Delete application](docs/openshift-delete-app.png)

2. Read the openshift.yaml to understand the Kubernetes resources and what they do.
3. Modify `openshift.yaml` to use your Quay.io image.
4. Login to the cluster from a CLI by copying the **Login Command**
   
    ![OpenShift Login command](docs/openshift-login-command.png)

5. You should now be able to do several commands using the OpenShift CLI:
   * `oc get pods` to see your pods that are running
   * `oc get deployments` will show you the deployment resource that OpenShift created for us in the previous step
   * `oc get svc` will show the services
   * `oc get routes` will show the routes & the urls from which you can connect
   * `oc apply -f openshift.yaml` will deploy your application
6. Automate this deployment process by modifying the `cd.yaml` pipeline.
   * Provide the correct secrets to login to the OpenShift cluster from your repository.
   * Make sure that this pipeline only runs from the dev branch.
   * The token that we used before is only valid for 24 hours, so we usually don't use this for automation but it's OK for this task.
7. Make a change in the openshift.yaml (update the replicas) and let the pipeline apply the changes.
8. Validate the changes via the UI or via the CLI.

---

### **Task 5: Updating the Application Code**
#### **Problem Statement**
Our `/workshop` endpoint lacks customization and fun elements.

#### **Solution Approach**
We'll update the application to:
* Add your **name** to the list of participants.
* Introduce a new field: `SweaterScore` that holds a numeric value of 1-10 on the presentators sweaters.
* Provide **validation** on this range.
* Set the **default SweaterScore via an environment variable**.

#### **Steps**
1. Modify the application code (`workshop.go`).
2. Validate that it works locally.
3. Release this update as **version 1.1.0**.

---

### **Task 6: Creating a New Environment (TST)**
#### **Problem Statement**
Right now, we only have a **DEV** environment. We need a **TST** (Test) environment.

#### **Solution Approach**
We'll create a separate **TST branch** and adjust configurations accordingly.

#### **Steps**
1. Create a new branch `tst` from `dev`.
2. Modify the `openshift.yaml` to:
    * Use different names for **Deployment, Service, and Route** by using `tst-` instead of `dev-`.
      * Normally we would use "Namespaces" to avoid naming conflicts, but the OpenShift Sandbox does not allow us to create multiple namespaces due to security risks.
    * Ensure it deploys only released versions (not latest dev builds).
3. Update the pipeline to trigger on the **TST branch**.
4. Deploy version **1.0.0**, then later **update to 1.1.0**.

---

## **Congratulations! 🎉**
You now have a fully functional **CI/CD pipeline**, supporting automated builds, releases, and deployments!

The next step is to now improve upon this, so that you as a developer can focus solely on producing the code. Choose some of the extra tasks to do, in both repositories are different tasks related to their type.

### **Extras**
Try improving your setup with:
* Custom Docker tags
  * Do custom tags for your docker builds, right now it takes the branch name but try something as followed:
    * On `main` branch, should be tagged as such `<name>:latest`.
    * On other branches, should be tagged as such `<name>:dev-<commitHash>` with the hash being a substring of 8 characters.
* Automatic deployment triggers
  * The `dev` Openshift environment should always the use the `:latest` image that is created when a developer pushes code to main and the build succeeds.
  * Create a trigger in your pipeline that automatically restarts the deployment on Openshift.
* Lightweight container images
  * Right now we're using the default Golang Docker image but we can speed up the build process by picking a smaller base image.
* Automated tests in Docker
  * Write a Go test & add tests execution to the multi-stage Dockerfile.
* Adding a metrics endpoint
  * Add a metrics endpoint to your application that will give you insights into the application performance.
* Implementing Helm Charts
  * [Helm charts](https://helm.sh/) are the most popular way of managing Kubernetes resources, because you can use advanced templating features.
  * Helm charts will replace our current kubectl apply and deploy method.
* Configuring Horizontal Pod Autoscaling
  * Add a Horizontal Pod Autoscaler that based on amount of load (requests) to the service will automatically spin up extra pods.
  * Simulate a load with Postman or just a bash while loop with curl.

Happy coding! 🚀