#!/usr/bin/env python3
import streamlit as st
import numpy as np
import pandas as pd
import matplotlib.pyplot as plt

tab1, tab2, tab3 = st.tabs(["Attack ⚡️", "Targets 📑", "Campaign 🧙🏽"])

with tab1:
    st.header("Attack ⚡️")
    st.subheader("Preview of Attack")
    st.image("https://purplesec.us/wp-content/uploads/2020/08/example-of-a-smishing-attack.png")

with tab2:
    st.header("Targets 📑")
    st.subheader("List of Targets / Employees")
    df = pd.DataFrame({
        "Name" : ["Antonio H. Blackman", "Veronica T. Branstetter", "William J. Williams"],
        "Number":["315-497-6643", "619-229-4726", "810-315-0389"],
        "Tags":["CEO", "President", "Intern"]
    })
    st.dataframe(df)

with tab3:
    st.header("Campaign 🧙🏽")

    st.subheader("Smishing Results")
    # Bar Graph
    bc = pd.DataFrame.from_dict({
        "Sent":[3,0,0],
        "Clicked":[0,2,0],
        "Login":[0,0,1]
    }, 
    orient='index', 
    columns=["Sent", "Clicked", "login"])

    st.bar_chart(bc)

    st.subheader("Assets affected")
    st.markdown('[![](https://mermaid.ink/img/pako:eNpVzbEOgjAQBuBXaW66JvAC3RQYnDTRscuFHtAIrTlKjCG8u9Xg4E1__v9LboU2OgYD3Rif7UCS1K22QeU7IFbNWWtVlqU6Il6EZ-84JK138F0qxFNILOG_rhGb0PvALFpDARPLRN7lT-uHWUgDT2zB5OhI7hZs2LJbHo4SN86nKGA6GmcugJYUr6_Qgkmy8A_VnnqhaVfbGzl_Pi0)](https://mermaid.live/edit#pako:eNpVzbEOgjAQBuBXaW66JvAC3RQYnDTRscuFHtAIrTlKjCG8u9Xg4E1__v9LboU2OgYD3Rif7UCS1K22QeU7IFbNWWtVlqU6Il6EZ-84JK138F0qxFNILOG_rhGb0PvALFpDARPLRN7lT-uHWUgDT2zB5OhI7hZs2LJbHo4SN86nKGA6GmcugJYUr6_Qgkmy8A_VnnqhaVfbGzl_Pi0)')



st.write("Send feedback: https://forms.gle/LvxyvjR1wEpQdCh27")
