/// <reference types="cypress" />

context('Setup Process', () => {

  it('should be not be setup yet', () => {
    cy.request(`/api`).then((response) => {
      expect(response.body).to.have.property('setup', false)
    })
  })

  it('should setup Statping with SQLite', () => {
    cy.visit('/setup', {failOnStatusCode: false})
    cy.get('#db_connection').select('sqlite')
    cy.get('#project').clear().type('Demo Tester')
    cy.get('#description').clear().type('This is a test from Crypress!')
    cy.get('#domain').clear().type('http://localhost:8888')
    cy.get('#username').clear().type('admin')
    cy.get('#password').clear().type('admin')
    cy.get('#password_confirm').clear().type('admin')
    cy.get('button[type="submit"]').click()

    cy.get('#title').should('contain', 'Demo Tester')
    cy.get('#description').should('contain', 'This is a test from Crypress!')
  })

    it('should have sample data', () => {
        cy.visit('/')
      cy.get('#title').should('contain', 'Demo Tester')
      cy.get('#description').should('contain', 'This is a test from Crypress!')
        cy.get('.card').should('have.length', 5)
        cy.get('.group_header').should('have.length', 2)
    })

  it('should be completely setup', () => {
    cy.request(`/api`).then((response) => {
      expect(response.body).to.have.property('setup', true)
      expect(response.body).to.have.property('domain', 'http://localhost:8888')
    })
  })

  it('should be able to Login', () => {
    cy.visit('/login')
    cy.get('#username').clear().type('admin')
    cy.get('#password').clear().type('admin')
    cy.get('button[type="submit"]').click()

    cy.get('.navbar-brand').should('contain', 'Statping')
    cy.getCookies()

    cy.getCookies().should('have.length', 1)

    cy.request(`/api`).then((response) => {
      expect(response.body).to.have.property('admin', true)
      expect(response.body).to.have.property('logged_in', true)
    })
  })

})
