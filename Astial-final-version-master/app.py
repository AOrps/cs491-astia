from apps.login import *
import requests
from streamlit_option_menu import option_menu
from streamlit_lottie import st_lottie
from multiapp import MultiApp
from apps import login, signup


st.set_page_config(page_title="Astia a Smishing Platform", page_icon=":fishing_pole_and_fish:", layout="wide")


# Use local CSS
def local_css(file_name):
    with open(file_name) as f:
        st.markdown(f"<style>{f.read()}</style>", unsafe_allow_html=True)


local_css("style/style.css")


# -------- Lottie Asset -------- #
def load_lottie(url):
    r = requests.get(url)
    if r.status_code != 200:
        return None
    return r.json()


animation = load_lottie("https://assets4.lottiefiles.com/packages/lf20_mcvtkrvc.json")
animation4 = load_lottie("https://assets5.lottiefiles.com/packages/lf20_isbiybfh.json")
animation5 = load_lottie("https://assets4.lottiefiles.com/packages/lf20_joexwr4o.json")
service3 = load_lottie("https://assets7.lottiefiles.com/packages/lf20_zdtukd5q.json")




# ----------- Sidebar ----------- #
with st.sidebar:
    selected = option_menu(
        menu_title="Astia Dashboard",
        options=["Home", "Services", "Pricing", "Support"],
        icons=["house", "card-checklist", "tags", "headset"],
        menu_icon=[""],
        default_index=0,
        styles={
            "nav-link": {
                "font-size": "20px",
                "--hover-color": "rgb(3, 171, 112)"
            },
            "nav-link-selected": {"background-color": "rgb(3, 171, 112)", "color": "white"}
        }
    )

# ------------- Home Page ------------------ #
if selected == "Home":
    string = "<h3 style = 'color:white; background-color:rgb(3, 171, 112); margin-top:-15px; font-style:italic;width: 100%;'>Astia<sub><i>Smishing Analysis Platform</i></sub></h3>"
    st.write(string, unsafe_allow_html=True)
    with st.sidebar:
        st.warning("Login or Signup by clicking Services for Launching Smishing Campaign")
    with st.container():
        left_column, right_column = st.columns(2)
        with left_column:
            st.write("###")
            st.write("###")
            st.subheader("This is where Cybersecurity meets Data Science")
            st.write("###")
            string = "Astia is a cybersecurity startup, specifically within user training awareness. The goal is to " \
                     "design a smishing platform that is capable of launching smishing campaigns and then measure " \
                     "analytics and metrics that has been generated by a campaign to make it easy to see patterns to " \
                     "mitigate real-world cyber attacks. "
            st.write(string, unsafe_allow_html=True)

        with right_column:
            st.write("###")
            st.write("###")
            st_lottie(animation, height=300, key="hacking")
    st.write("###")
    with st.container():
        st.subheader("Data Analytics & Visualization")
        st.write("###")
        string = "It's important to not only show that we are capable of launching campaigns but we want to also offer a" \
                 "myriad of data generated by a campaign to help security teams take a data-driven approach to locking down their company."
        st.write(string, unsafe_allow_html=True)

# ------------- Service Page ------------------ #
if selected == "Services":

    with st.container():
        string = "<h3 style = 'color:white; background-color:rgb(3, 171, 112); margin-top:-15px; font-style:italic;width: 100%;'>Astia<sub><i>Smishing Analysis Platform</i></sub></h3>"
        st.write(string, unsafe_allow_html=True)
        st.write("##")

        left_column, right_column = st.columns(2)
        with left_column:
            st.subheader("Smishing Campaign")
            st.write("While your employees are your greatest asset, they can also be your biggest threat regarding "
                     "cyber security. Between phishing, social engineering attacks, malware distribution, "
                     "password theft and breaches, there has never been a greater need for non-technical employees to "
                     "maintain a cyber culture.")

        with right_column:
            st_lottie(service3, height=200, key="service3")
        with st.sidebar:
            app = MultiApp()
            app.add_app("Login", login.app)
            app.add_app("SignUp", signup.app)
            app.run()
        login.success()

# ------------- Support Page ------------------ #
if selected == "Support":
    string = "<h3 style = 'color:white; background-color:rgb(3, 171, 112); margin-top:-15px; font-style:italic;width: 100%;'>Astia<sub><i>Smishing Analysis Platform</i></sub></h3>"
    st.write(string, unsafe_allow_html=True)
    st.write("##")
    st.write("##")
    with st.container():
        st.subheader("Need Help? Astia is one step away from you!")
        st.write("##")
        contact_form = """
        <form action="https://formsubmit.co/naveedanwar808@gmail.com" method="POST">
        <input type="hidden" value="false" name="_captcha">
     <input type="text" placeholder="Your name" name="name" required>
     <input type="email" placeholder="Your email" name="email" required>
     <textarea name="message" placeholder="Your message" required></textarea>
     <button type="submit">Send</button>
    </form>
        """
        left_column, right_column = st.columns(2)
        with left_column:
            st.markdown(contact_form, unsafe_allow_html=True)
        with right_column:
            st_lottie(animation4, height=300, key="service")

# ------------- Pricing Page ------------------ #
if selected == "Pricing":
    string = "<h3 style = 'color:white; background-color:rgb(3, 171, 112); margin-top:-15px; font-style:italic;width: 100%;'>Astia<sub><i>Smishing Analysis Platform</i></sub></h3>"
    st.write(string, unsafe_allow_html=True)
    with st.container():
        st.write("##")
        st.write("##")
        left_column, right_column = st.columns(2)
        with left_column:
            st.write("")
        with right_column:
            st.write("")
